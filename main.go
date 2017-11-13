package main

import (
	"flag"
	"time"
	"fmt"
	"os"
	"net/http"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"strings"
)

// some structs representing the menus retrieved from an API
type Menu struct {
	Number int16 `json:"number"`
	Year   int16 `json:"year"`
	Days   []Day `json:"days"`
}

type Day struct {
	Date   string `json:"date"`
	Dishes []Dish `json:"dishes"`
}

type Dish struct {
	Name  string `json:"name"`
	Price interface{} `json:"price"`
}

// date format for parsing
const format = "2006-01-02"

// dictionary for mapping short names to full names of locations
var locations = map[string]string{
	"mg": "mensa-garching",
	"ma": "mensa-arcisstrasse",
	"sg": "stubistro-grosshadern",
	"fmi": "fmi-bistro",
}

// API urls
var apis = map[string]string {
    "mensa-garching": "https://srehwald.github.io/eat-api/",
    "mensa-arcisstrasse": "https://srehwald.github.io/eat-api/",
    "stubistro-grosshadern": "https://srehwald.github.io/eat-api/",
    "fmi-bistro": "https://srehwald.github.io/eat-api/",
}

var currentDate = time.Now()

/*
Retrieve the menu for a given location and date. Returns a Menu struct representing all menus for a specific week.
 */
func getMenu(location string, date time.Time) (*Menu, error) {
	// convert year to string
	year := strconv.Itoa(date.Year())
	// get week number
	_, w := date.ISOWeek()
	// convert week number to string
	week := strconv.Itoa(w)
	// add leading zero if necessary
	if len(week) < 2 {
		week = "0" + week
	}

    api := apis[location]

	// build url
	url := api + location + "/" + year + "/" + week + ".json"
	// make GET request
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// parse body into struct
	s := new(Menu)
	unmarshalErr := json.Unmarshal(body, &s)

	return s, unmarshalErr
}

/*
Stringify the dishes of a Day struct. Returns the resulting string as well as the length of the longest line.
 */
func dishesToString(day Day) (dishesStr string, maxLength int) {
    for _, dish := range day.Dishes {
        var dishStr string

        // check if price is of type float64
        price, ok := dish.Price.(float64)
        if ok {
            // if price is float, convert float to string
            dishStr = dish.Name + ": " + strconv.FormatFloat(price, 'f', -1, 64) + "€"
        } else {
            /*
            if price is not float, it is most likely a string not containing the price, but something
            like "Self Service"
             */
            priceStr, ok := dish.Price.(string)
            if ok {
                dishStr = dish.Name + ": " + priceStr
            } else {
                // if price is neither float nor string, it is not available
                dishStr = dish.Name + ": Not available"
            }
        }
        if len(dishStr) > maxLength {
            maxLength = len(dishStr)
        }
        dishesStr += "\n" + dishStr
    }

    return dishesStr, maxLength
}

/*
Find a specific Day given a date
 */
func findDay(date string, days []Day) (day Day, found bool) {
	// loop through all days and compare dates
    for _, d := range days {
        if d.Date == date {
            return d, true
        }
    }

    return day, false
}

/*
Print usage
 */
func showUsage() {
    fmt.Println("usage: eat [-options] <location>")
    fmt.Println("Options:")
	fmt.Println("    -h, --help \tshow usage")
    fmt.Println("    -d DATE \tdate of the menu (format: yyyy-mm-dd; default: current date)")
    fmt.Println("Locations: <full name> (<short name>)")
    for k,v := range locations {
        fmt.Println("    "+ v + " (" + k + ")")
    }
}

/*
MAIN
 */
func main() {
    // override default usage menu
    flag.Usage = showUsage

    // define date option
	dateArg := flag.String("d", currentDate.Format(format), "date of the menu")

    // parse args and options
	flag.Parse()

    // parse date option
	date, err := time.Parse(format, *dateArg)
	if err != nil {
		fmt.Println("Error: Cannot parse date '" + *dateArg + "'. (Required format: yyyy-mm-dd)")
		os.Exit(1)
	}

    // setting a location is mandatory
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Error: Missing location")
		os.Exit(1)
	}

    // the first arg will be considered as location
	location := args[0]
    if Contains(location, Keys(locations)) {
        // get full location name
        location = locations[location]
    } else if !Contains(location, Values(locations)) {
		fmt.Println("Error: Location '" + location + "' not found.")
		os.Exit(1)
	}

	message := "Menu for '" + location + "' on '" + date.Format(format) + "':"
	fmt.Println(message)

    // get menu for given date and location
	menu, err := getMenu(location, date)
	if err != nil {
		fmt.Println("Error: Could not get menu.")
		os.Exit(1)
	}

	// find correct day given the date
	day, foundDay := findDay(*dateArg, menu.Days)
	if !foundDay {
		fmt.Println("Error: Could not find menu for your date '" + *dateArg + "'.")
		os.Exit(0)
	}

	hlineLength := len(message)
	dishesStr, maxLength := dishesToString(day)
    if maxLength > hlineLength {
        hlineLength = maxLength
    }

	// create and print horizontal line
	hline := strings.Repeat("-", hlineLength)
	fmt.Print(hline)
	// print dishes
	fmt.Println(dishesStr)
}

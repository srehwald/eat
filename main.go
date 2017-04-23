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
)

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

const format = "2006-01-02"
const api = "https://srehwald.github.io/stwm-mensa-api/"

var currentDate = time.Now()

// TODO make more generic for different locations from different APIs
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

	// build url
	url := api + location + "/" + year + "/" + week + ".json"
	// make GET request
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	// read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	// parse body into struct
	var s = new(Menu)
	unmarshalErr := json.Unmarshal(body, &s)

	return s, unmarshalErr
}

func main() {
	// TODO show available location in help menu

	dateArg := flag.String("date", currentDate.Format(format), "date of the menu")

	flag.Parse()

	date, err := time.Parse(format, *dateArg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var args = flag.Args()
	if len(args) < 1 {
		fmt.Println("Error: missing location")
		os.Exit(1)
	}
	var location = args[0]

	fmt.Println("Menu for '" + location + "' on '" + date.Format(format) + "':")

	// TODO error message, if location is wrong
	menu, err := getMenu(location, date)
	if err != nil {
		fmt.Println("Error: Could not get menu.")
		os.Exit(1)
	}

	// find correct day given the date
	var day Day
	var foundDay = false
	for _, d := range menu.Days {
		if d.Date == *dateArg {
			day = d
			foundDay = true
			break
		}
	}

	if !foundDay {
		fmt.Println("Could not find menu for your date '" + *dateArg + "'.")
		os.Exit(0)
	}

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

		fmt.Println(dishStr)
	}
}

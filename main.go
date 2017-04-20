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
    Year int16 `json:"year"`
    Days []Day `json:"days"`
}

type Day struct {
    Date string `json:"date"`
    Dishes []Dish `json:"dishes"`
}

type Dish struct {
    Name string `json:"name"`
    Price float64 `json:"price"`
}

// TODO better format notation
const format = "2006-01-02"
const api = "https://srehwald.github.io/stwm-mensa-api/"
var currentDate = time.Now()


func getMenu(location string, date time.Time) (*Menu, error) {
    year := strconv.Itoa(date.Year())
    _, w := date.ISOWeek()
    week := strconv.Itoa(w)

    // TODO leading zeros
    url := api + location + "/" + year + "/" + week + ".json"
    res, err := http.Get(url)
    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }

    var s = new(Menu)
    unmarshalErr := json.Unmarshal(body, &s)

    return s, unmarshalErr
}


func main() {

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

    fmt.Println("Menu for '"+ location+ "' on '" + date.Format(format) + "':")
    menu, err := getMenu(location, date)

    // TODO get correct date
    day := menu.Days[0]

    for _, dish := range day.Dishes {
        fmt.Println(dish.Name + ": " + strconv.FormatFloat(dish.Price, 'f', -1, 64) + "€")
    }
}

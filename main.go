package main

import (
    "flag"
    "time"
    "fmt"
    "os"
)

// TODO better format notation
const format = "2006-01-02"
var currentDate = time.Now()

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
}

# eat

[![Build Status](https://travis-ci.org/srehwald/eat.svg?branch=master)](https://travis-ci.org/srehwald/eat)

Command line tool for getting the daily menu of different locations written in Go.

## Build
Clone this repository and run `go build` within the directory. 
```
$ git clone https://github.com/srehwald/eat.git
$ cd eat/
$ go build
```

## Available Locations
By now, the following locations are available:
- Mensa Garching
- Mensa Arcisstraße
- StuBistro Großhadern

Call `eat -h` in order to see the correct names you need to use within the tool.

## Usage

Check out the usage menu:
```
usage: eat [-options] <location>
Options:
    -h, --help  show usage
    -d DATE     date of the menu (format: yyyy-mm-dd; default: current date)
Locations: <full name> (<short name>)
    mensa-garching (mg)
    mensa-arcisstrasse (ma)
    stubistro-grosshadern (sg)

```

#### Examples
The following examples assume that you added the `eat` executable to your `PATH`. Otherwise you might need to call it with `./eat` (in case you are in the same directory as the executable).

Get the current menu for mensa-garching (You can also use the short name 'mg'):
```
$ eat mensa-garching
Menu for 'mensa-garching' on '2017-04-28':
-------------------------------------------------------------------
Bauerneintopf: 1€
Pfannkuchen mit Apfelmus: 1.55€
Gedünsteter Kabeljau (MSC) auf Buchweizen-Rote-Beete-Risotto: 4€

$ eat mg
Menu for 'mensa-garching' on '2017-04-28':
-------------------------------------------------------------------
Bauerneintopf: 1€
Pfannkuchen mit Apfelmus: 1.55€
Gedünsteter Kabeljau (MSC) auf Buchweizen-Rote-Beete-Risotto: 4€

```

Get the current menu for mensa-arcisstrasse on 2017-04-27:
```
$ eat eat -d 2017-04-27 mensa-arcisstrasse
Menu for 'mensa-arcisstrasse' on '2017-04-27':
------------------------------------------------------------------
Reistopf mit Chinagemüse: 1€
Quinoa-Salat mit Nüssen, Gurke, Rucola und Mandeldressing: 2.4€
Grünes Hähnchen-Curry: 3.2€
Reistopf mit Chinagemüse (2): Self-Service
Pikante Nudelpfanne mit Schafskäse und Rindfleisch: Self-Service
Pasta mit Bärlauchpesto: Self-Service

```
# Go Figure

## Description
Go Figure prints beautiful ASCII art from text.
It supports [FIGlet](http://www.figlet.org/) files,
and most of its features.

This package was inspired by the Ruby gem [artii](https://github.com/miketierney/artii),
but built from scratch.

## Installation
`go get github.com/common-nighthawk/go-figure`

## Basic Example
```go
package main

import("github.com/common-nighthawk/go-figure")

func main() {
  myFigure := figure.NewFigure("Hello World", "")
  myFigure.Print()
}
```

```txt
  _   _          _   _          __        __                 _       _ 
 | | | |   ___  | | | |   ___   \ \      / /   ___    _ __  | |   __| |
 | |_| |  / _ \ | | | |  / _ \   \ \ /\ / /   / _ \  | '__| | |  / _` |
 |  _  | |  __/ | | | | | (_) |   \ V  V /   | (_) | | |    | | | (_| |
 |_| |_|  \___| |_| |_|  \___/     \_/\_/     \___/  |_|    |_|  \__,_|
                                                                       

```

## Documentation
### Create a Figure
The constructor take two arguments, the text and font name.
If passed an empty string for the font name, a default is provided.
That is, these are both valid--

`myFigure := figure.NewFigure("Foo Bar", "")`

`myFigure := figure.NewFigure("Foo Bar", "alphabet")`

Please note that font names are case sensitive
and only standard ASCII characters are supported
(character codes 32-127).

### Methods: stdout
#### Print()
The most basic, and common, method is func Print.
A figure responds to Print(), and will write the output to the terminal.
There is no return value.

`myFigure.Print()`

But if you really want to have a good time,
explore the other methods below.

#### Blink(duration, timeOn, timeOff int)
A figure responds to the func Blink, taking three arguments.
`duration` is the total time the banner will display, in milliseconds.
`timeOn` is the length of time the text will blink on (also in ms).
`timeOff` is the length of time the text will blink off (ms).
For an even blink, set `timeOff` to -1
(same as setting `timeOff` to the value of `timeOn`).
There is no return value.

`myFigure.Blink(5000, 1000, 500)`

`myFigure.Blink(5000, 1000, -1)`

#### Scroll(duration, stillness int, direction string)
A figure responds to the func Scroll, taking three arguments.
`duration` is the total time the banner will display, in milliseconds.
`stillness` is the length of time the text will not move (also in ms).
So, the lower the stillness the faster the scroll speed.
`direction` can either be "right" or "left" (case insensitive).
The direction will be left if an invalid option is passed.
There is no return value.

`myFigure.Scroll(5000, 200, "right")`

`myFigure.Scroll(5000, 100, "left")`

#### Dance(duration, freeze int)
A figure responds to the func Dance, taking two arguments.
`duration` is the total time the banner will display, in milliseconds.
`freeze` is the length of time between dance moves (also in ms).
So, the lower the freeze the faster the dancing.
There is no return value.

`myFigure.Scroll(5000, 800)`

### Methods: Writers
#### Write(w io.Writer, fig figure)
Unlike the above methods that operate on a figure value,
func Write is a function that takes two arguments.
`w` is a value that implements all the methods in the io.Writer interface.
`fig` is the figure that will be written.

`figure.Write(w, myFigure)`

This method is can be useful in adding a nifty banner to a web page--
`
func landingPage(w http.ResponseWriter, r *http.Request) {
  figure.Write(w, figure.NewFigure("Hello World", "alphabet"))
`

### Methods: Misc
#### Slicify() ([]string)
If you wish to do something outside of the created methods,
you can always grab the internal slice.
This should give you a good start to build anything
with the ASCII art, if manually.

A figure responds to the func Slicify,
and will return the slice of string.

`myFigure.Slicify()`
returns
`["FFFF           BBBB         ",
  "F              B   B        ",
  "FFF  ooo ooo   BBBB   aa rrr",
  "F    o o o o   B   B a a r  ",
  "F    ooo ooo   BBBB  aaa r  "]`


## More Examples
`figure.NewFigure("Go-Figure", "isometric1").Print()`

```
      ___           ___           ___                       ___           ___           ___           ___     
     /\  \         /\  \         /\  \          ___        /\  \         /\__\         /\  \         /\  \    
    /::\  \       /::\  \       /::\  \        /\  \      /::\  \       /:/  /        /::\  \       /::\  \   
   /:/\:\  \     /:/\:\  \     /:/\:\  \       \:\  \    /:/\:\  \     /:/  /        /:/\:\  \     /:/\:\  \  
  /:/  \:\  \   /:/  \:\  \   /::\~\:\  \      /::\__\  /:/  \:\  \   /:/  /  ___   /::\~\:\  \   /::\~\:\  \ 
 /:/__/_\:\__\ /:/__/ \:\__\ /:/\:\ \:\__\  __/:/\/__/ /:/__/_\:\__\ /:/__/  /\__\ /:/\:\ \:\__\ /:/\:\ \:\__\
 \:\  /\ \/__/ \:\  \ /:/  / \/__\:\ \/__/ /\/:/  /    \:\  /\ \/__/ \:\  \ /:/  / \/_|::\/:/  / \:\~\:\ \/__/
  \:\ \:\__\    \:\  /:/  /       \:\__\   \::/__/      \:\ \:\__\    \:\  /:/  /     |:|::/  /   \:\ \:\__\  
   \:\/:/  /     \:\/:/  /         \/__/    \:\__\       \:\/:/  /     \:\/:/  /      |:|\/__/     \:\ \/__/  
    \::/  /       \::/  /                    \/__/        \::/  /       \::/  /       |:|  |        \:\__\    
     \/__/         \/__/                                   \/__/         \/__/         \|__|         \/__/    
```

`figure.NewFigure("Foo Bar Pop", "smkeyboard").Print()`

```
 ____  ____  ____  ____  ____  ____  ____  ____  ____ 
||F ||||o ||||o ||||B ||||a ||||r ||||P ||||o ||||p ||
||__||||__||||__||||__||||__||||__||||__||||__||||__||
|/__\||/__\||/__\||/__\||/__\||/__\||/__\||/__\||/__\|
```

`figure.NewFigure("Keep Your Eyes On Me", "rectangles").Print()`

```
                                                                                          
 _____                 __ __                 _____                 _____       _____      
|  |  | ___  ___  ___ |  |  | ___  _ _  ___ |   __| _ _  ___  ___ |     | ___ |     | ___ 
|    -|| -_|| -_|| . ||_   _|| . || | ||  _||   __|| | || -_||_ -||  |  ||   || | | || -_|
|__|__||___||___||  _|  |_|  |___||___||_|  |_____||_  ||___||___||_____||_|_||_|_|_||___|
                 |_|                               |___|                                  
```

`figure.NewFigure("ABCDEFGHIJ", "eftichess").Print()`

```
#########         #########   ___   #########         #########                           
##[`'`']#  \`~'/  ##'\v/`##  /\*/\  ##|`+'|##  '\v/`  ##\`~'/##  [`'`']   '\v/`    \`~'/  
###|  |##  (o o)  ##(o 0)## /(o o)\ ##(o o)##  (o 0)  ##(o o)##   |  |    (o 0)    (o o)  
###|__|##   \ / \ ###(_)###   (_)   ###(_)###   (_)   ###\ / \#   |__|     (_)      \ / \ 
#########    "    #########         #########         ####"####                      "    
```

`figure.NewFigure("It's been waiting for you", "doom").Blink(10000, 500, -1)`

![blink](url "blink")

`figure.NewFigure("I mean, I could...", "doom").Scroll(10000, 400, "right")`
`figure.NewFigure("But why would I want to?", "doom").Scroll(10000, 400, "left"`

![scroll](url "scroll")

`figure.NewFigure("Give your reasons", "doom").Dance(10000, 400)`

![dance](url "dance")


## List of Fonts
* 3-d
* 3x5
* 5lineoblique
* acrobatic
* alligator
* alligator2
* alphabet
* avatar
* banner
* banner3-D
* banner3
* banner4
* barbwire
* basic
* bell
* big
* bigchief
* binary
* block
* bubble
* bulbhead
* calgphy2
* caligraphy
* catwalk
* chunky
* coinstak
* colossal
* computer
* contessa
* contrast
* cosmic
* cosmike
* cricket
* cursive
* cyberlarge
* cybermedium
* cybersmall
* diamond
* digital
* doh
* doom
* dotmatrix
* drpepper
* eftichess
* eftifont
* eftipiti
* eftirobot
* eftitalic
* eftiwall
* eftiwater
* epic
* fender
* fourtops
* fuzzy
* goofy
* gothic
* graffiti
* hollywood
* invita
* isometric1
* isometric2
* isometric3
* isometric4
* italic
* ivrit
* jazmine
* jerusalem
* katakana
* kban
* larry3d
* lcd
* lean
* letters
* linux
* lockergnome
* madrid
* marquee
* maxfour
* mike
* mini
* mirror
* mnemonic
* morse
* moscow
* nancyj-fancy
* nancyj-underlined
* nancyj
* nipples
* ntgreek
* o8
* ogre
* pawp
* peaks
* pebbles
* pepper
* poison
* puffy
* pyramid
* rectangles
* relief
* relief2
* rev
* roman
* rot13
* rounded
* rowancap
* rozzo
* runic
* runyc
* sblood
* script
* serifcap
* shadow
* short
* slant
* slide
* slscript
* small
* smisome1
* smkeyboard
* smscript
* smshadow
* smslant
* smtengwar
* speed
* stampatello
* standard
* starwars
* stellar
* stop
* straight
* tanja
* tengwar
* term
* thick
* thin
* threepoint
* ticks
* ticksslant
* tinker-toy
* tombstone
* trek
* tsalagi
* twopoint
* univers
* usaflag
* wavy
* weird

## Contributing
Because this projects is small, we can dispense with formalities.
Submit a pull request, open an issue, request a change.
All good!

## Wanna Say Thanks?
GitHub stars are helpful.
Most importantly, they help with discoverability.
Projects with more stars are displayed higher
in seach results when people are looking for packages.
Also--they make contributors feel good :)

If you are feeling especially generous,
give a shout to [@cmmn_nighthawk](https://twitter.com/cmmn_nighthawk).

## TODO
* Add proper support for spaces
* More animations
* Implement graceful line-wrapping and smushing

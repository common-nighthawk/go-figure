# Go Figure

## Description
Go Figure prints beautiful ASCII art from text.
It supports [FIGlet](http://www.figlet.org/) files,
and most of its features.

This package was inspired by the Ruby gem [artii](https://github.com/miketierney/artii),
but built from scratch.

## Installation
`go get github.com/common-nighthawk/go-figure`

## Example
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
#### Create a Figure
The constructor take two arguments, the text and font name.
If passed an empty string for the font name, a default is provided.
That is, these are both valid--

`myFigure := figure.NewFigure("Foo Bar", "")`

`myFigure := figure.NewFigure("Foo Bar", "alphabet")`

Please note that font names are case sensitive
and only standard ASCII characters are supported
(character codes 32-127).

#### Print It
A figure response to func Print(). It provides no return value.

`myFigure.Print()`

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

## TODO
* Add proper support for spaces
* Do not display below baseline when empty
* New feature: animation
* Implement graceful line-wrapping and smushing

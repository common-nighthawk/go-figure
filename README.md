# Go Figure

## Description
Prints ASCII Art from text.  Currently only support [a-zA-Z0-9] characters.

Inspired by the [artii](https://github.com/miketierney/artii) Ruby gem.
But not nearly as fully-features.

The font name is [Alphabet](http://www.figlet.org/fontdb_example.cgi?font=alphabet.flf) created for FIGlet.

## Installation
`go get github.com/common-nighthawk/go-figure`

## Example
```go
package main

import("github.com/common-nighthawk/go-figure")

func main() {
  figure.Print("Hello World")
}
```

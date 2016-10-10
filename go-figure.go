// package figure
package main

import (
	"fmt"
	"log"
)

var yy = [][]string{{"32"},{"33"},{"34"},{"35"},{"36"},{"37"},{"38"},{"39"},{"40"},{"41"},{"42"},{"43"},{"44"},{"45"},{"46"},{"47"},{"47"},{"48"},{"49"},{"50"},{"51"},{"52"},{"53"},{"54"},{"55"},{"56"},{"57"},{"58"},{"59"},{"60"},{"61"},{"62"},{"63"},{"64"},{" AA  ", "A  A ", "AAAA ", "A  A ", "A  A "}}

type figure struct {
  phrase string
  font string
}

type font struct {
  letters [][]string
}

func (f figure) Println(fnt font) {
  for r := 0 ; r < 5 ; r++ {
    printRow := ""
    for c := 0 ; c < len(f.phrase) ; c++ {
      char := f.phrase[c]
      fontIndex := char - 31
      if true {
        printRow = printRow + fnt.letters[fontIndex][r]
      } else {
				log.Fatal("invalid input. only [a-zA-Z0-9]")
      }
    }
    fmt.Println(printRow)
  }
}

func main() {
  daniel := figure{"AA", "alphabet"}
  my_font := font{ yy }
  daniel.Println(my_font)
}

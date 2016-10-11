// package figure
package main

import (
	"fmt"
	"log"
  "os"
  "bufio"
)

// var yy = [][]string{{"32"},{"33"},{"34"},{"35"},{"36"},{"37"},{"38"},{"39"},{"40"},{"41"},{"42"},{"43"},{"44"},{"45"},{"46"},{"47"},{"47"},{"48"},{"49"},{"50"},{"51"},{"52"},{"53"},{"54"},{"55"},{"56"},{"57"},{"58"},{"59"},{"60"},{"61"},{"62"},{"63"},{"64"},{" AA  ", "A  A ", "AAAA ", "A  A ", "A  A ", "    ", "    "},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}}

type figure struct {
  phrase string
  font string
}

type font struct {
  letters [][]string
  height int
}

func NewFont(name string) font {
  var fnt font
  fnt.height = 7
  fnt.letters = [][]string{{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}}

  fnt.letters[0] = []string{ "   ", "   ", "   ", "   ", "   ", "   ", "   ", "   " }
  file, _ := os.Open("fonts/alphabet.flf")
  defer file.Close()
  scanner := bufio.NewScanner(file)

  counter := 0
  for scanner.Scan() {
    text := scanner.Text()
    cutLength := 1

    if counter > 0 {
      if text[len(text)-2:] == "@@" {
        cutLength = 2
      }
      fnt.letters[counter] = append(fnt.letters[counter], text[:len(text)-cutLength])
    }

    if text == "@@@@@@@@" {
      counter = 1
    } else if len(text) > 2 && text[len(text)-2:] == "@@" {
      counter += 1
    }
  }

  return fnt
}

func (f figure) Println(fnt font) {
  for r := 0 ; r < fnt.height ; r++ {
    printRow := ""
    for c := 0 ; c < len(f.phrase) ; c++ {
      char := f.phrase[c]
      fontIndex := char - 32
      if true {
        printRow = printRow + fnt.letters[fontIndex][r]
      } else {
				log.Fatal("invalid input.")
      }
    }
    fmt.Println(printRow)
  }
}

func main() {
  daniel := figure{"DANIEL allen DEUTSCH", "alphabet"}
  my_font := NewFont(daniel.font)
  daniel.Println(my_font)
}

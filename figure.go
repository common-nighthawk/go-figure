package figure

import (
	"fmt"
	"log"
  "strings"
)

const ascii_offset = 32
const first_ascii = ' '
const last_ascii = '~'

type figure struct {
  phrase string
  font font
}

func NewFigure(phrase string, fontName string) figure {
  return figure{phrase, NewFont(fontName)}
}

func scrubText(text string, char byte) string {
  return strings.Replace(text, string(char), " ", -1)
}

func (figure figure) Print() {
  var font = figure.font
  for r := 0 ; r < font.height ; r++ {
    printRow := ""
    for c := 0 ; c < len(figure.phrase) ; c++ {
      char := figure.phrase[c]
      if char >= first_ascii && char <= last_ascii {
        fontIndex := char - ascii_offset
        printRow = printRow + scrubText(font.letters[fontIndex][r], font.hardBlank)
      } else {
				log.Fatal("invalid input.")
      }
    }
    fmt.Println(printRow)
  }
}

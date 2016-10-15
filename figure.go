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
  return figure{phrase, newFont(fontName)}
}

func scrub(text string, char byte) string {
  return strings.Replace(text, string(char), " ", -1)
}

func (figure figure) Print() {
  for r := 0 ; r < figure.font.height ; r++ {
    printRow := ""
    for c := 0 ; c < len(figure.phrase) ; c++ {
      char := figure.phrase[c]
      if char >= first_ascii && char <= last_ascii {
        fontIndex := char - ascii_offset
        charRowText := scrub(figure.font.letters[fontIndex][r], figure.font.hardBlank)
        printRow += charRowText
      } else {
				log.Fatal("invalid input.")
      }
    }
    fmt.Println(printRow)
  }
}

package figure

import (
	"fmt"
	"log"
)

type Figure struct {
  Phrase string
  Font Font
}

func NewFigure(phrase string, fontName string) Figure {
  if fontName == "" {
    fontName = "alphabet"
  }
  return Figure{ phrase, NewFont(fontName) }
}

func (fig Figure) Print() {
  for r := 0 ; r < fig.Font.Height ; r++ {
    printRow := ""
    for c := 0 ; c < len(fig.Phrase) ; c++ {
      char := fig.Phrase[c]
      if char >= ' ' && char <= '~' {
        fontIndex := char - 32
        printRow = printRow + fig.Font.Letters[fontIndex][r]
      } else {
				log.Fatal("invalid input.")
      }
    }
    fmt.Println(printRow)
  }
}

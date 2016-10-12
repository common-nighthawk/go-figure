package figure

import (
	"fmt"
	"log"
  "os"
  "bufio"
  "strconv"
)

type Figure struct {
  Phrase string
  Font string
}

type Font struct {
  Letters [][]string
  Height int
}

func NewFont(name string) Font {
  var fnt Font

  file, _ := os.Open("/Users/Daniel/Documents/go-workspace/src/go-figure/fonts/alphabet.flf")
  defer file.Close()
  scanner := bufio.NewScanner(file)

  fnt.Letters = append(fnt.Letters, []string{})

  counter := 0
  for scanner.Scan() {
    text := scanner.Text()
    if fnt.Height == 0 && text[:4] == "flf2" {
      fnt.Height, _ = strconv.Atoi(string(text[5]))
    }


    fnt.Letters = append(fnt.Letters, []string{})
    cutLength := 1

    if counter > 0 {
      if text[len(text)-2:] == "@@" {
        cutLength = 2
      }
      fnt.Letters[counter] = append(fnt.Letters[counter], text[:len(text)-cutLength])
    }

    if text == "@@@@@@@@" {
      counter = 1
    } else if len(text) > 2 && text[len(text)-2:] == "@@" {
      counter += 1
    }
  }
  for i := 0 ; i < fnt.Height ; i ++ {
    fnt.Letters[0] = append(fnt.Letters[0], "   ")
  }

  return fnt
}

func (f Figure) Println(fnt Font) {
  for r := 0 ; r < fnt.Height ; r++ {
    printRow := ""
    for c := 0 ; c < len(f.Phrase) ; c++ {
      char := f.Phrase[c]
      fontIndex := char - 32
      if true {
        printRow = printRow + fnt.Letters[fontIndex][r]
      } else {
				log.Fatal("invalid input.")
      }
    }
    fmt.Println(printRow)
  }
}

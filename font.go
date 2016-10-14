package figure

import (
	"fmt"
  "os"
  "bufio"
  "strconv"
  "regexp"
)

type Font struct {
  Letters [][]string
  Height int
}

func NewFont(name string) Font {
  var fnt Font

  file_path := fmt.Sprintf("../figure/fonts/%s.flf", name)
  file, _ := os.Open(file_path)
  defer file.Close()
  scanner := bufio.NewScanner(file)

  fnt.Letters = append(fnt.Letters, []string{})

  counter := 0
  for scanner.Scan() {
    text := scanner.Text()
    if fnt.Height == 0 && text[:4] == "flf2" {
      r, _ := regexp.Compile(`\d+`)
      matches := r. FindAllString(text, -1)
      fnt.Height, _ = strconv.Atoi(matches[1])
    }


    fnt.Letters = append(fnt.Letters, []string{})
    cutLength := 1

    if counter > 0 {
      if len(text) > 1 && (text[len(text)-2:] == "@@" || text[len(text)-2:] == "##") {
        cutLength = 2
      }
      if len(text) > 1 {
        fnt.Letters[counter] = append(fnt.Letters[counter], text[:len(text)-cutLength])
      }
    }

    if len(text) > 1 && (text[len(text)-2:] == "@@" || text[len(text)-2:] == "##") {
      counter ++
    }
  }
  for i := 0 ; i < fnt.Height ; i ++ {
    fnt.Letters[0] = append(fnt.Letters[0], "   ")
  }

  return fnt
}

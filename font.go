package figure

import (
	"fmt"
  "os"
  "bufio"
  "strconv"
  "regexp"
)

type font struct {
  letters [][]string
  height int
  hardBlank byte
}

// func setHeight(font font, metadata string) font {
//   r, _ := regexp.Compile(`\d+`)
//   matches := r. FindAllString(text, -1)
//   font.height, _ = strconv.Atoi(matches[1])
// }

// func setHardBlank(font font, metadata string) font {
//   r, _ := regexp.Compile(`f\S+\s`)
//   matches := r.FindAllString(text, -1)
//   mymatch := matches[0]
//   font.hardBlank = ' '
//   if mymatch[len(mymatch)-2] != 'a' && mymatch[len(mymatch)-2] != '2' {
//     font.hardBlank = mymatch[len(mymatch)-2]
//   }
// }

func NewFont(name string) font {
  var font font

  //TODO change path
  file_path := fmt.Sprintf("../figure/fonts/%s.flf", name)
  file, _ := os.Open(file_path)
  defer file.Close()
  scanner := bufio.NewScanner(file)

  font.letters = append(font.letters, []string{})
  counter := 0
  for scanner.Scan() {
    text := scanner.Text()
    if font.height == 0 && text[:4] == "flf2" {
      // setHeight(font, text)
      // setHardBlank(font, text)

      r, _ := regexp.Compile(`\d+`)
      matches := r. FindAllString(text, -1)
      font.height, _ = strconv.Atoi(matches[1])

      r2, _ := regexp.Compile(`f\S+\s`)
      matches2 := r2.FindAllString(text, -1)
      mymatch := matches2[0]
      font.hardBlank = ' '
      if mymatch[len(mymatch)-2] != 'a' && mymatch[len(mymatch)-2] != '2' {
        font.hardBlank = mymatch[len(mymatch)-2]
      }
    }


    font.letters = append(font.letters, []string{})
    cutLength := 1

    if counter > 0 {
      if len(text) > 1 && (text[len(text)-2:] == "@@" || text[len(text)-2:] == "##" || text[len(text)-2:] == "$$") {
        cutLength = 2
      }
      if len(text) > 1 {
        font.letters[counter] = append(font.letters[counter], text[:len(text)-cutLength])
      }
    }

    if len(text) > 1 && (text[len(text)-2:] == "@@" || text[len(text)-2:] == "##" || text[len(text)-2:] == "$$") {
      counter ++
    }
  }
  for i := 0 ; i < font.height ; i ++ {
    font.letters[0] = append(font.letters[0], "   ")
  }

  return font
}

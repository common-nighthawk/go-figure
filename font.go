package figure

import (
	"fmt"
  "os"
  "bufio"
  "strconv"
  "regexp"
)

var hardBlanksBlacklist = [2]byte{'a', '2'}

type font struct {
  name string
  hardBlank byte
  height int
  letters [][]string
}

func NewFont(name string) font {
  font := font{name: name}
  growLetters(&font)
  setDefaults(&font)

  //TODO change path
  file_path := fmt.Sprintf("../figure/fonts/%s.flf", font.name)
  file, _ := os.Open(file_path)
  defer file.Close()

  letterIndex := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text, cutLength, indexInc := scanner.Text(), 1, 0
    growLetters(&font)
    if font.height == 0 && text[:4] == "flf2" {
      setHeight(&font, text)
      setHardBlank(&font, text)
    }
    if lastCharLine(text, font.height) {
      indexInc = 1
    }
    if lastCharLine(text, font.height) && font.height > 1 {
      cutLength = 2
    }
    if letterIndex > 0 {
      appendText := ""
      if len(text) > 1 {
        appendText = text[:len(text)-cutLength]
      }
      font.letters[letterIndex] = append(font.letters[letterIndex], appendText)
    }
    letterIndex += indexInc
  }
  return font
}

func setDefaults (font *font) {
  font.hardBlank = ' '
  font.letters[0] = make([]string, 100, 100)
  //TODO: MAKE SURE FILE flf EXITS FOR NAME
  if len(font.name) < 1 {
    font.name = "standard"
  }
}

func setHeight(font *font, metadata string) {
  r, _ := regexp.Compile(`\d+`)
  matches := r. FindAllString(metadata, -1)
  font.height, _ = strconv.Atoi(matches[1])
}

func setHardBlank(font *font, metadata string) {
  r, _ := regexp.Compile(`f\S+\s`)
  match := r.FindAllString(metadata, -1)[0]
  possibleHardBlank := match[len(match)-2]
  if possibleHardBlank != hardBlanksBlacklist[0] && possibleHardBlank != hardBlanksBlacklist[1] {
    font.hardBlank = possibleHardBlank
  }
}

func growLetters(font *font) {
  font.letters = append(font.letters, []string{})
}

func lastCharLine(text string, height int) bool {
  if height > 1 {
    lastChars := "  "
    if len(text) > 1 {
      lastChars = text[len(text)-2:]
    }
    return lastChars == "@@" || lastChars == "##" || lastChars == "$$"
  } else {
    lastChar := " "
    if len(text) > 0 {
      lastChar = text[len(text)-1:]
    }
    return lastChar == "@" || lastChar == "#" || lastChar == "$"
  }
}

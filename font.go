package figure

import (
  "bufio"
  "os"
)

type font struct {
  name string
  hardBlank byte
  height int
  letters [][]string
}

func newFont(name string) (font font) {
  setAttributes(&font, name)
  file := getFile(font.name)
  defer file.Close()
  setLetters(&font, file)
  return font
}

func setAttributes(font *font, name string) {
  font.name = name
  growLetters(font)
  font.letters[0] = make([]string, 100, 100)
  if len(name) < 1 {
    font.name = "standard"
  }
}

func setLetters(font *font, file *os.File) {
  letterIndex := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text, cutLength, indexInc := scanner.Text(), 1, 0
    growLetters(font)
    if font.height == 0 && text[:4] == "flf2" {
      font.height = getHeight(text)
      font.hardBlank = getHardBlank(text)
    }
    if lastCharLine(text, font.height) {
      indexInc = 1
      if font.height > 1 { cutLength = 2 }
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
}

func growLetters(font *font) {
  font.letters = append(font.letters, []string{})
}

package figure

import (
  "bufio"
)

type font struct {
  name string
  height int
  baseline int
  hardblank byte
  reverse bool
  letters [][]string
}

func newFont(name string) (font font) {
  setName(&font, name)
  file := getFile(font.name)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  setAttributes(&font, scanner)
  setLetters(&font, scanner)
  return font
}

func setName(font *font, name string) {
  font.name = name
  if len(name) < 1 {
    font.name = "standard"
  }
}

func setAttributes(font *font, scanner *bufio.Scanner) {
  for scanner.Scan() {
    text := scanner.Text()
    if text[:4] == signature {
      font.height = getHeight(text)
      font.baseline = getBaseline(text)
      font.hardblank = getHardblank(text)
      font.reverse = getReverse(text)
      break
    }
  }
}

func setLetters(font *font, scanner *bufio.Scanner) {
  extendLetters(font)
  font.letters[0] = make([]string, 100, 100) //TODO: set spaces from flf
  letterIndex := 0
  for scanner.Scan() {
    text, cutLength, letterIndexInc := scanner.Text(), 1, 0
    if lastCharLine(text, font.height) {
      extendLetters(font)
      letterIndexInc = 1
      if font.height > 1 { cutLength = 2 }
    }
    if letterIndex > 0 {
      appendText := ""
      if len(text) > 1 {
        appendText = text[:len(text)-cutLength]
      }
      font.letters[letterIndex] = append(font.letters[letterIndex], appendText)
    }
    letterIndex += letterIndexInc
  }
}

func extendLetters(font *font) {
  font.letters = append(font.letters, []string{})
}

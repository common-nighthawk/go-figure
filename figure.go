package figure

import (
  "fmt"
  "io"
  "log"
  "strings"
  "time"
)

const ascii_offset = 32
const first_ascii = ' '
const last_ascii = '~'

type figure struct {
  phrase string
  font font
}

func NewFigure(phrase string, fontName string) figure {
  font := newFont(fontName)
  if font.reverse {
    phrase = reverse(phrase)
  }
  return figure{phrase, font}
}

func (figure figure) Print() {
  for _, printRow := range figure.Slicify() {
    fmt.Println(printRow)
  }
}

func (figure figure) Scroll(duration int, freq int) {
  endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
  figure.phrase = figure.phrase + "   "
  fmt.Println("\033[H\033[2J")
  for time.Now().Before(endTime) {
    chars := strings.Split(figure.phrase, "")
    chars = append(chars, chars[0])
    chars = chars[1:len(chars)]
    newS := strings.Join(chars, "")
    figure.phrase = newS
    figure.Print()
    time.Sleep(time.Duration(freq) * time.Millisecond)
    fmt.Println("\033[H\033[2J")
  }
}

func (figure figure) Blink(duration int, freq int) {
  endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
  fmt.Println("\033[H\033[2J")
  for time.Now().Before(endTime) {
    figure.Print()
    time.Sleep(time.Duration(freq) * time.Millisecond)
    fmt.Println("\033[H\033[2J")
    time.Sleep(time.Duration(freq) * time.Millisecond)
  }
}

func Write(w io.Writer, figure figure) {
  for _, printRow := range figure.Slicify() {
    fmt.Fprintf(w, "%v\n", printRow)
  }
}

func (figure figure) Slicify() (rows []string) {
  for r := 0 ; r < figure.font.height ; r++ {
    printRow := ""
    for c := 0 ; c < len(figure.phrase) ; c++ {
      char := figure.phrase[c]
      if char >= first_ascii && char <= last_ascii {
        fontIndex := char - ascii_offset
        charRowText := scrub(figure.font.letters[fontIndex][r], figure.font.hardblank)
        printRow += charRowText
      } else {
        log.Fatal("invalid input.")
      }
    }
    if r < figure.font.baseline || len(strings.TrimSpace(printRow)) > 0 {
      rows = append(rows, printRow)
    }
  }
  return rows
}

func scrub(text string, char byte) string {
  return strings.Replace(text, string(char), " ", -1)
}

func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

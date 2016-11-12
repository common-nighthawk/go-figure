package figure

import (
  "fmt"
  "io"
  "strings"
  "time"
)

//stdout
func (figure figure) Print() {
  for _, printRow := range figure.Slicify() {
    fmt.Println(printRow)
  }
}

func (figure figure) Scroll(duration, stillness int, direction string) {
  endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
  figure.phrase = figure.phrase + "   "
  clearScreen()
  for time.Now().Before(endTime) {
    var shiftedPhrase string
    chars := []byte(figure.phrase)
    if strings.HasPrefix(strings.ToLower(direction), "r") {
      shiftedPhrase = string(append(chars[len(chars)-1:], chars[0:len(chars)-1]...))
    } else {
      shiftedPhrase = string(append(chars[1:len(chars)], chars[0]))
    }
    figure.phrase = shiftedPhrase
    figure.Print()
    sleep(stillness)
    clearScreen()
  }
}

func (figure figure) Blink(duration, timeOn, timeOff int) {
  if timeOff < 0 { timeOff = timeOn }
  endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
  clearScreen()
  for time.Now().Before(endTime) {
    figure.Print()
    sleep(timeOn)
    clearScreen()
    sleep(timeOff)
  }
}

func (myFig figure) Dance(duration, freeze int) {
  endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
  spacer := strings.Repeat("", 3)
  clearScreen()
  figures := []figure{NewFigure("", myFig.font.name), NewFigure("", myFig.font.name)}
  for i, c := range myFig.phrase {
    appenders := []string{spacer, spacer}
    appenders[i%2] = string(c)
    for f, _ := range figures {
      figures[f].phrase = figures[f].phrase + appenders[f]
    }
  }
  for p := 0; time.Now().Before(endTime); p ^= 1 {
    figures[p].Print()
    figures[1-p].Print()
    sleep(freeze)
    clearScreen()
  }
}

func clearScreen() {
  fmt.Print("\033[H\033[2J")
}

func sleep(milliseconds int) {
  time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

//writers
func Write(w io.Writer, figure figure) {
  for _, printRow := range figure.Slicify() {
    fmt.Fprintf(w, "%v\n", printRow)
  }
}

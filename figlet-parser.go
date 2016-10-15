package figure

import (
  "fmt"
  "log"
  "os"
  "regexp"
  "strconv"
  "strings"
)

var hardBlanksBlacklist = [2]byte{'a', '2'}

func getFile(name string) (file *os.File) {
  //TODO change path
  file_path := fmt.Sprintf("../figure/fonts/%s.flf", name)
  file, err := os.Open(file_path)
  if err != nil {
    log.Fatal("invalid font name.")
  }
  return file
}

func getHeight(metadata string) int {
  r, _ := regexp.Compile(`\d+`)
  matches := r. FindAllString(metadata, -1)
  height, _ := strconv.Atoi(matches[1])
  return height
}

func getHardBlank(metadata string) byte {
  r, _ := regexp.Compile(`f\S+\s`)
  match := r.FindAllString(metadata, -1)[0]
  hardBlank := match[len(match)-2]
  if hardBlank != hardBlanksBlacklist[0] && hardBlank != hardBlanksBlacklist[1] {
    return hardBlank
  } else {
    return ' '
  }
}

func lastCharLine(text string, height int) bool {
  endOfLine, length := "  ", 2
  if height == 1 && len(text) > 0 {
    length = 1
  }
  if len(text) >= length {
    endOfLine = text[len(text)-length:]
  }
  return endOfLine == strings.Repeat("@", length) || endOfLine == strings.Repeat("#", length) || endOfLine == strings.Repeat("$", length)
}

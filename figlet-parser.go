package figure

import (
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

const signature = "flf2"
const reverseFlag = "1"
var charDelimiters = [3]string{"@", "#", "$"}
var hardBlanksBlacklist = [2]byte{'a', '2'}

func getFile(name string) (file *os.File) {
  file_path := fmt.Sprintf("../figure/fonts/%s.flf", name) //TODO change path
  file, err := os.Open(file_path)
  if err != nil {
    log.Fatal("invalid font name.")
  }
  return file
}

func getHeight(metadata string) int {
  datum := strings.Fields(metadata)[1]
  height, _ := strconv.Atoi(datum)
  return height
}

func getHardBlank(metadata string) byte {
  datum := strings.Fields(metadata)[0]
  hardBlank := datum[len(datum)-1]
  if hardBlank == hardBlanksBlacklist[0] || hardBlank == hardBlanksBlacklist[1] {
    return ' '
  } else {
    return hardBlank
  }
}

func getReverse(metadata string) bool {
  data := strings.Fields(metadata)
  return len(data) > 6 && data[6] == reverseFlag
}

func lastCharLine(text string, height int) bool {
  endOfLine, length := "  ", 2
  if height == 1 && len(text) > 0 {
    length = 1
  }
  if len(text) >= length {
    endOfLine = text[len(text)-length:]
  }
  return endOfLine == strings.Repeat(charDelimiters[0], length) ||
           endOfLine == strings.Repeat(charDelimiters[1], length) ||
           endOfLine == strings.Repeat(charDelimiters[2], length)
}

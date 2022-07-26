package utils

import(
  "io/ioutil"
  "log"
)

func GetContentFromFile(filePath string) string {
  content, err := ioutil.ReadFile(filePath)
  if err != nil {
      log.Fatal(err)
  }

  // Convert []byte to string and print to screen
  text := string(content)
  return text
}

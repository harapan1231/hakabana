package main

import (
  "fmt"
  "os"
  "io/ioutil"
)

func main() {
  b, err := ioutil.ReadFile(os.Args[1])
  if err != nil {
    panic(err)
  }
  fmt.Printf("Hello, world.\n%s\n", b)
}

package main

import (
    "fmt"
    "io/ioutil"
    "bytes"
)

func main() {
    body, _ := ioutil.ReadFile("textfile_1000kb.txt")
    result := bytes.Title(body)
    fmt.Printf("%s", result)
}
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	body, _ := ioutil.ReadFile("textfile_1000kb.txt")
	lines := bytes.Title(body)
	fmt.Printf("%s", lines)
	ioutil.WriteFile("textfile_caps_go.txt", lines, 0644)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	// common
	reset  = "\033[0m" // auto reset the rest of text to default color
	normal = 0
	bold   = 1 // increase this value if you want bolder text
	// special
	dim       = 2
	underline = 4
	blink     = 5
	reverse   = 7
	hidden    = 8
	// color
	black       = 30 // default = 39
	red         = 31
	green       = 32
	yellow      = 33
	blue        = 34
	purple      = 35 // purple = magenta
	cyan        = 36
	lightGray   = 37
	darkGray    = 90
	lightRed    = 91
	lightGreen  = 92
	lightYellow = 93
	lightBlue   = 94
	lightPurple = 95
	lightCyan   = 96
	white       = 97
)

type msg struct {
	Message string `json:"msg"` //= {"msg":"Message Body content"}
}

func main() {

	port := os.Getenv("PORT") // port string
	if port == "" {
		port = "8080"
	}

	fmt.Println("Go HTTPServer running on http://localhost:" + port + " [Ctrl+C to stop] \n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fname := r.URL.Path[1:]
		http.ServeFile(w, r, r.URL.Path[1:])
		t := time.Now()

		fmt.Printf("%s - %s file: %s\n", t, r.RemoteAddr, fname)
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {

		if err := json.NewEncoder(w).Encode(&msg{Message: "Hello Go Lang JSON"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":"+port, nil)
}

func clr(colorCode int, content string) string {

	return "\033[" + strconv.Itoa(normal) + ";" + strconv.Itoa(colorCode) + "m" + content + reset
	//return "\033[" + strconv.Itoa(colorCode) + "m" + content + reset

}

func formatBytes(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

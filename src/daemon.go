package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	// Bet Sure to issue go get github.com/sevlyar/go-daemon first to pull down github
	"github.com/sevlyar/go-daemon"
)

// To terminate the daemon use:
//  kill `cat daemon.pid`
func main() {
	cntxt := &daemon.Context{
		PidFileName: "daemon.pid",
		PidFilePerm: 0644,
		LogFileName: "daemon.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon daemon]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	serveHTTP()
}

func serveHTTP() {
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
	log.Print("- - - - - - - - - - - - - - -")
	log.Print("HTTP Server started at: 127.0.0.1:8080")

}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, "go-daemon: %q", html.EscapeString(r.URL.Path))
}

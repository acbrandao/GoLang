///////////////////////////////////////////////////////////
// HTTPSSE.go : Simple HTTP Server sent  Events server
///////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	//to generate unique id
)

const (
	// TIMEOUT in seconds
	TIMEOUT      = 30
	DEBUG   bool = true
)

func main() {

	port := os.Getenv("PORT") // port string
	if port == "" {
		port = "8080"
	}

	fmt.Println("Go HTTP Server Sent Events running on http://localhost:" + port + " [Ctrl+C to stop] \n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fname := r.URL.Path[1:]
		http.ServeFile(w, r, r.URL.Path[1:])
		t := time.Now()

		fmt.Printf("%s - %s file: %s\n", t, r.RemoteAddr, fname)
	})

	http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		count := 0
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			//return nil, ErrStreamingNotSupported
		}

		// Set the headers related to event streaming.
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		//ticker := time.NewTicker(1000 * time.Millisecond)
		for count < 10 {
			count++
			//now write out the SSE message here
			msg_id := uniqid()
			fmt.Fprintf(w, "id: %s\n", msg_id)
			fmt.Fprintf(w, "event: %s\n", "timeout")
			fmt.Fprintf(w, "data: %ds\n\n", count)
			flusher.Flush()

			time.Sleep(500 * time.Millisecond)

			fmt.Println("Ticker count ", count)

		}

	})

	http.ListenAndServe(":"+port, nil)
}

func uniqid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	// uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	uuid := fmt.Sprintf("%x", b[0:4])

	return uuid
}

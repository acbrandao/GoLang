package main

import (
    "log"
	"net/http"
	"os"
    "fmt"
)

func main() {
	port := os.Getenv("PORT")  // port string
	if port == "" {
		port = "443"
	}

	 rootPath:="./static"

    http.Handle("/", http.FileServer(http.Dir(rootPath)))

	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
	fmt.Printf("HTTP/S Server https://localhost:%s  Root Path = %s",port,rootPath)
}
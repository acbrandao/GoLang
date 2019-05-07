package main
import(
    "os"
    "fmt"
    "net/http"
	"encoding/json"
	"log"
	"io/ioutil"
)

type msg struct {
    Message string `json:"msg"` //= {"msg":"Message Body content"}
}

func main() {
	port := os.Getenv("PORT")  // port string
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on http://localhost:",port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {

        if err := json.NewEncoder(w).Encode( &msg{ Message: "Hello Go Lang JSON"} ); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
	})
	
    http.ListenAndServe(":"+port, nil)
}

func readJsonFile(fn string)string {

	stream,err :=ioutil.ReadFile(fn)
	if err!=nil{
		log.Fatal(err)
	}

	data:=string(stream)
	return data
}
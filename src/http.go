package main
import(
    "os"
    "fmt"
	"net/http"
	"encoding/json"
	"time"
	"syscall"
)

type msg struct {
    Message string `json:"msg"` //= {"msg":"Message Body content"}
}

func main() {

	port := os.Getenv("PORT")  // port string
	if port == "" {
		port = "8080"
	}


	fmt.Printf("Server running on http://localhost:%s  (Ctrl+C to stop) \n",port)

	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fname :=r.URL.Path[1:]
		http.ServeFile(w, r, r.URL.Path[1:])
		t := time.Now()
   
		fmt.Printf("%s - %s file: %s\n",t,r.RemoteAddr,fname )
    })

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {

        if err := json.NewEncoder(w).Encode( &msg{ Message: "Hello Go Lang JSON"} ); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/sys", func(w http.ResponseWriter, r *http.Request) {
		var stat syscall.Statfs_t
		wd, err := os.Getwd()

		if err !=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		syscall.Statfs(wd, &stat)

		// Available blocks * size per block = available space in bytes
		var disk_space=stat.Bavail * uint64(stat.Bsize)

		fmt.Fprintf(w, "Disk Space: %s",formatBytes(int64(disk_space)  ))
	})
	
	
    http.ListenAndServe(":"+port, nil)
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

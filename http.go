// Serve current directory on HTTP
// Defaults to port 8080
package main
import (
    "log"
    "net/http"
    "os"
)

func logHandler(fs http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("[%v][%v]\n", r.RemoteAddr, r.URL)
        fs.ServeHTTP(w, r)
    }
}

func main() {
    addr := ":8080"
    if len(os.Args)>1 {
        addr = os.Args[1]
    }
    log.Println("Listening on "+addr)
    fileServer := http.FileServer(http.Dir("."))
    http.ListenAndServe(addr, logHandler(fileServer))
}


// Serve current directory on HTTP
// Defaults to port 8080
package main
import (
    "log"
    "net/http"
    "os"
)
func main() {
    addr := ":8080"
    if len(os.Args)>1 {
        addr = os.Args[1]
    }
    log.Println("Listening on "+addr)
    panic(http.ListenAndServe(addr, http.FileServer(http.Dir("."))))
}


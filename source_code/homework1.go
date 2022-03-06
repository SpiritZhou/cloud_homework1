package main
import (
    "fmt"
	"log"
    "os"
	"net/http"
    "net"
)


func main() {
    fmt.Println("start");
    os.Setenv("VERSION", "1");
    http.HandleFunc("/healthz", healthz);
    http.HandleFunc("/", index);
    http.ListenAndServe(":8080", nil)

}

func healthz(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
}

func index(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("VERSION", os.Getenv("VERSION"))
    
    for key, element := range r.Header {
        for _, content := range element {
            w.Header().Set(key, content)
        }
    }

    ip := r.Header.Get("X-Real-IP")
    if ip == "" {
        ip,_,_ = net.SplitHostPort(r.RemoteAddr)
    }
    log.Printf("Response Status:", 200)
    log.Printf("Client IP:", ip)
}
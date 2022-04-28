package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rand.Seed(time.Now().UnixNano())
}

const addr =":4567"
const pacScript = `
function FindProxyForURL(url, host)
{
    var tunnel = 'SOCKS 192.168.1.2:8889';
    return tunnel;
}
`

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",ServeHTTP)
	log.Println("listen at:",addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
}

func ServeHTTP(w http.ResponseWriter,r *http.Request){
	w.Write([]byte(pacScript))
	w.Header().Add("Content-Type","application/x-ns-proxy-autoconfig")
	w.WriteHeader(200)
}
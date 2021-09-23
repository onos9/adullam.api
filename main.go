package main

import (
	"flag"
	"log"
	"net/http"

	"beznet/adullam/server"
)

var (
	version = "1.0"
	addr    = flag.String("addr", ":8000", "http service address")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	server.Init()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/"+r.URL.Path[1:])
	})

	http.HandleFunc("/ws", server.WebsocketHandler)

	go func() {
	   server.Run()
	}()
	
	log.Println("WebRTC Signaling Server NetCam. version="+version)
	log.Println("running on http://localhost"+*addr+" (Press Ctrl+C quit)",)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
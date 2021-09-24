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
	maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
	maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/"+r.URL.Path[1:])
	})

	// Create the job queue.
	jobQueue := make(chan server.Job, *maxQueueSize)
	
	// Start the dispatcher.
	dispatcher := server.NewDispatcher(jobQueue, *maxWorkers)
	dispatcher.Run()

	server.HandleApiRequest()
	
	log.Println("Adullam api server started. version="+version)
	log.Println("running on http://localhost"+*addr+" (Press Ctrl+C quit)",)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func payloadHandler(){}

func request(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

		// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg map[string]interface{}
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	m, ok := msg["name"].(string)
	if !ok {
		m = "default"
	}

	d, ok := msg["data"].(map[string]interface{})
	if !ok {
		d = make(map[string]interface{})
	}

	// Create Job and push the work onto the jobQueue.
	job := Job{method: m, data: d}
	queue <- job

	// Render success.
	w.WriteHeader(http.StatusCreated)
}

func HandleApiRequest(){
	http.HandleFunc("/api/v1/auth", authenticate)
	http.Handle("/api/v1/seminary", isAuthorized(request))
}
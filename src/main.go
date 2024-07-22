package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func Fnlong(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Set the content type to text/event-stream
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	vars := mux.Vars(r)
	name := vars["name"]
	i := 0
	count, _ := strconv.Atoi(name)
	for {
		i++
		fmt.Fprintf(w, "all: %v  . current: %v\n", count, i)
		flusher.Flush() // Flush the buffer to the client
		if i > count {
			break

		}
		time.Sleep(time.Second)

	}
}

func FnlongWaitAll(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Set the content type to text/event-stream
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	vars := mux.Vars(r)
	name := vars["name"]
	i := 0
	count, _ := strconv.Atoi(name)
	for {
		i++
		// fmt.Fprintf(w, "all: %v  . current: %v\n", count, i)
		// flusher.Flush() // Flush the buffer to the client
		if i > count {
			break

		}
		time.Sleep(time.Second)

	}
	fmt.Fprintf(w, "all: %v  . current: %v\n", count, i)
	flusher.Flush() // Flush the buffer to the client

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloHandler)
	r.HandleFunc("/long/{name}", Fnlong).Methods("GET")
	r.HandleFunc("/longconn/{name}", Fnlong).Methods("GET")
	r.HandleFunc("/wait/{name}", FnlongWaitAll).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Starting server on :80")
	http.ListenAndServe(":80", nil)
}

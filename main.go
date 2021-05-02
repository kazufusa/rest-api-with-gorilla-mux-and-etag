package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Clock struct {
	Unix   int64  `json:"unix"`
	String string `json:"string"`
	json   []byte `json:"-"`
}

func NewClock() (clock *Clock, err error) {
	now := time.Now().Truncate(time.Minute)
	clock = &Clock{
		Unix:   now.Unix(),
		String: now.String(),
	}
	clock.json, err = json.Marshal(clock)
	if err != nil {
		return nil, err
	}
	return
}

func (c *Clock) etag() string {
	return fmt.Sprintf("%x", sha1.Sum(c.json))
}

func ClockHandler(w http.ResponseWriter, r *http.Request) {
	clock, err := NewClock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	etag := clock.etag()

	if val := r.Header.Get("If-None-Match"); val == etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("ETag", etag)
	_, err = w.Write(clock.json)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/clock", ClockHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}

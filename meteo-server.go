/*
The server accepts meteo data from weather station as GET request.
It stores the data in memory and provides it to the clients as a key-value pairs.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Prevents concurrent map iteration and write
var mutex = sync.Mutex{}

// Meteo data from weather station
var meteoData = make(map[string]string)

func handleMeteo(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()

	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, "Could not parse query parameters", http.StatusBadRequest)
		return
	}

	for key := range values {
		meteoData[key] = values.Get(key)
	}

	if len(values) > 0 {
		meteoData["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	}

	for key, value := range meteoData {
		fmt.Fprintf(w, "\n%s=%s", key, value)
	}

	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", handleMeteo)

	fmt.Printf("Starting server on port 8080...\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

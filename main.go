// Basic Go Server Side Prototype
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	DISTANCE = "200km"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	User     string   `json:"user"`
	Message  string   `json:"message"`
	Location Location `json:"location"`
}

func main() {
	fmt.Println("service starting")
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// route "/post" handler. Decode request body in JSON and write the message field to response.
// w: std writer object to write in response
// r: a pointer of request field
// Post request. Use r.body to get all the jason field
func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one post request.")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}

// route "/search" handler. Get fields from URL and formed as Post Structure.
// w: std writer object to write in response
// r: a pointer of request field
// Get request. Use string coverter (strenv) and r.URL to get related field
func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	ran := DISTANCE
	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}
	fmt.Println("range is", ran)
	p := &Post{
		User:    "1111",
		Message: "The 100 places you mush go",
		Location: Location{
			Lat: lat,
			Lon: lon,
		},
	}
	jsonString, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonString)
}

//18215036/Ahmad Fawwaz Zuhdi
package main

import (
	"log"
	"net/http"
	"encoding/json"

	"nautica/core"

	"github.com/gorilla/mux"
)

func CreateDiveSequence(w http.ResponseWriter, r *http.Request) {
	var sequence core.DiveSeq
	var unparsed string
	_ = json.NewDecoder(r.Body).Decode(&sequence)
	sequence.ID = (*r).RemoteAddr
	sequence, unparsed = core.Calculate(sequence)
	note := []string{unparsed}
	core.WriteLines(note, "./DiveLog")
	json.NewEncoder(w).Encode(sequence)
}

func main() {
	addr := ":8081"
	router := mux.NewRouter()
	router.HandleFunc("/dives/{id}", CreateDiveSequence).Methods("POST")
	log.Println("nautica running on localhost", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
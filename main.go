package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"nautica/core"

	"github.com/gorilla/mux"
)

func CreateDiveSequence(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var sequence core.DiveSeq
	var unparsed string
	_ = json.NewDecoder(r.Body).Decode(&sequence)
	sequence.ID = params["id"]
	sequence, unparsed = core.Calculate(sequence)
	note := []string{unparsed}
	core.WriteLines(note, "./DiveLog")
	json.NewEncoder(w).Encode(sequence)
}

/*
func GetDiveSequence(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(diveSeq)
}
*/

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dives/{id}", CreateDiveSequence).Methods("POST")
	//router.HandleFunc("/dives/{id}", GetDiveSequence).Methods("GET")
	fmt.Println("nautica running on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

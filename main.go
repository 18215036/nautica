package main

import {
	"fmt"
	"log"
	"net/http"

	"nautica/core"

	"github.com/gorilla/mux"
}

var diveLog []DiveSeq

func CreateDiveSequence(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var sequence DiveSeq
	_ = json.NewDecoder(r.Body).Decode(&sequence)
	sequence.ID = params["id"]
	Calculate(sequence)
	diveLog = append(diveLog, sequence)
	json.NewEncoder(w).Encode(diveLog)
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
	log.Fatal(http.ListenAndServe(":8081", router))
}
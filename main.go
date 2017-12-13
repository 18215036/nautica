<<<<<<< HEAD
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
=======
package main

import (
	"log"
	"net/http"
	"encoding/json"

	"nautica/core"

	"github.com/gorilla/mux"
)

func CreateDiveSequence(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var sequence core.DiveSeq
	var unparsed string
	_ = json.NewDecoder(r.Body).Decode(&sequence)
	//sequence.ID = params["id"]
	log.Println(sequence)
	log.Println(sequence.Dive1)
	log.Println(sequence.Dive2)
	log.Println(sequence.Dive3)
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
	addr := ":8081"
	router := mux.NewRouter()
	router.HandleFunc("/dives/{id}", CreateDiveSequence).Methods("POST")
	//router.HandleFunc("/dives/{id}", GetDiveSequence).Methods("GET")
	
	log.Println("nautica running on localhost", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
>>>>>>> a0515de94cffc247399381688e840fcfac72213b

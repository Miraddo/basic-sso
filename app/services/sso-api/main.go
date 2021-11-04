package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/miraddo/basic-sso/app/tooling/database"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		os.Exit(1)
	}

	res, err := newUser.Insert()

	if err != nil {
		log.Fatalln(err)
	}

	if res == true {
		w.WriteHeader(http.StatusCreated)
		log.Printf("User Successfuly Added!")
	}

	w.WriteHeader(http.StatusBadRequest)
	log.Printf("User Already Exist!")

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/register", RegisterHandler)

	log.Println("Starting server on :5000")

	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}

// Package main is a basic project to handle SSO system.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/miraddo/basic-sso/app/tooling/database"
)

// RegisterHandler is got a request to register new user.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		os.Exit(1)
	}

	res, err := newUser.Insert()

	var status string

	switch {
	case err != nil:
		w.WriteHeader(http.StatusForbidden)
		status = err.Error()

	case res:
		w.WriteHeader(http.StatusCreated)
		status = "User Successfully Added!"

	default:
		w.WriteHeader(http.StatusForbidden)
		status = "User Already Exist!"
	}

	w.Write([]byte(status))
}

// main responsable to declare route and hande basic webservice.
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/register", RegisterHandler)

	log.Println("Starting server on :5000")

	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}

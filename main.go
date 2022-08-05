package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	serv "challenge/Services"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Starting")
	router := chi.NewRouter()
	router.Get("/api/GetTransaction", getHandler)
	router.Post("/api/PostTransaction", postHandler)
	router.Get("/api/GetTransaction/{Id}", GetTransaction)

	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Server is listining on port 8080")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	transactionlist := serv.GetAllTransaction()
	json.NewEncoder(w).Encode(transactionlist)
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	trans := &serv.Transaction{}
	json.NewDecoder(r.Body).Decode(trans)
	trans.CreatedAt = time.Now()
	trans.Id = rand.Intn(1000)
	w.Write([]byte("Done!"))
	w.WriteHeader(201)
}
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	Id := chi.URLParam(r, "Id")
	//check if the id is not null
	if Id == "" {
		http.Error(w, fmt.Sprintf("Not Found"), 404)
		return
	}

	//get transaction
	transaction := &serv.Transaction{}
	err, transaction := serv.GetTransaction(Id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Id: %s is not found", Id), 404)
		return
	}
	json.NewEncoder(w).Encode(transaction)
	w.WriteHeader(200)

}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/go/crud/models"
	"github.com/go/crud/storage"
	"github.com/gofiber/fiber/v2"
)

// func homeLink(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "Welcome home!")
// 	json.NewEncoded(w).Encode("Hi")
// }

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/book", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome home!")
	})
	
	log.Println("app is running")
	http.ListenAndServe(":4080", router)
}
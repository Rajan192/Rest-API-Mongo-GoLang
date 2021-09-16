package main

import (
	handler "USER/HandleFunc"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to This API Project of Cila Labs")
	r := mux.NewRouter()

	r.HandleFunc("/v1/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/v1/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/v1/user/{id}", handler.GetUserByID).Methods("GET")
	//r.HandleFunc("/user/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/v1/user/{id}", handler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

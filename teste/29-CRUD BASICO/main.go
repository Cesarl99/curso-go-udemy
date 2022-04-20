package main

import (
	"CRUD/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//CRUD - Create,Read,Update,Delete

//CREATE - POST
//READ - GET
//UPDATE - PUT
//DELETE - DELETE
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuarios).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuarios).Methods(http.MethodDelete)
	fmt.Println("Escutando na Porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

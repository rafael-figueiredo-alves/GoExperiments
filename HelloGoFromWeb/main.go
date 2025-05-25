package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"` //json: "firstname" é um Decorator para que o json seja em minúsculo o nome da propriedade
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {

	log.Print("Configurando endpoints")

	http.HandleFunc("/", func(Response http.ResponseWriter, Request *http.Request) {
		fmt.Fprintf(Response, "Hello, Go, from web endpoint: %s\n", Request.URL.Path)
	})

	log.Println("Iniciando segundo endpoint")

	http.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Exibindo página: %s\n", r.URL.Path)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var user = User{
			Firstname: "Rafael",
			Lastname:  "Alves",
			Age:       40,
		}

		json.NewEncoder(w).Encode(user)
	})

	log.Print("Iniciando servidor")

	log.Fatal(http.ListenAndServe(":80", nil))
}

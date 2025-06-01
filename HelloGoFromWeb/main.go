package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	Firstname string `json:"firstname"` //json: "firstname" é um Decorator para que o json seja em minúsculo o nome da propriedade
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {

	log.Print("Configurando endpoints")

	log.Println("Iniciando endpoints")

	http.HandleFunc("GET /clientes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Exibindo página: %s\n", r.URL.Path)
	})

	http.HandleFunc("GET /ping", func(Response http.ResponseWriter, Request *http.Request) {
		fmt.Fprintf(Response, "Pong")
	})

	http.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(fmt.Sprintf("HTTP method %q not allowed", r.Method)))
			return
		}

		var user = User{
			Firstname: "Rafael",
			Lastname:  "Alves",
			Age:       40,
		}

		json.NewEncoder(w).Encode(user)
	})

	http.HandleFunc("GET /", func(Response http.ResponseWriter, Request *http.Request) {
		if Request.URL.Path != "/" {
			http.NotFound(Response, Request)
			return
		}
		fmt.Fprintf(Response, "Hello, Go, from web endpoint: %s\n", Request.URL.Path)
	})

	log.Print("Iniciando servidor")

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Erro durante execução do servidor: %v\n", err)
		os.Exit(1)
	}
}

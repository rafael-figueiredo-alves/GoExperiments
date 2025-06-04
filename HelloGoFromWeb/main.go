package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
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

	http.HandleFunc("GET /clientes", GetClientes)

	http.HandleFunc("GET /testeheaders", GetTesteHeaders)

	http.HandleFunc("GET /ping", GetPing)

	http.HandleFunc("POST /user", PostUser)

	http.HandleFunc("GET /", GetHome)

	exibirIP()

	log.Print("Iniciando servidor")

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Erro durante execução do servidor: %v\n", err)
		os.Exit(1)
	}
}

func GetTesteHeaders(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, %s, %s", r.RemoteAddr, r.Method, r.UserAgent())
}

func GetClientes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Exibindo página: %s\n", r.URL.Path)
}

func GetPing(Response http.ResponseWriter, Request *http.Request) {
	fmt.Fprintf(Response, "Pong")
}

func PostUser(w http.ResponseWriter, r *http.Request) {
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
}

func GetHome(Response http.ResponseWriter, Request *http.Request) {
	if Request.URL.Path != "/" {
		http.NotFound(Response, Request)
		return
	}
	fmt.Fprintf(Response, "Hello, Go, from web endpoint: %s\n", Request.URL.Path)
}

func exibirIP() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Erro ao obter interfaces:", err)
		return
	}

	for _, iface := range interfaces {
		// Ignorar interfaces que estão desativadas ou que não suportam endereços
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Printf("Erro ao obter endereços da interface %s: %v\n", iface.Name, err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Exibir apenas endereços IPv4
			if ip != nil && ip.To4() != nil {
				fmt.Printf("Interface: %s, IP: %s\n", iface.Name, ip.String())
			}
		}
	}
}

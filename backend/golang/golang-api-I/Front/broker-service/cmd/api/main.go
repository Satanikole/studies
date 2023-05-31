package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Broker service iniciado na porta %s\n", webPort)

	// Definindo a porta http do server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Iniciando o servidor
	err := server.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

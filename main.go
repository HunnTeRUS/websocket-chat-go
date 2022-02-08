package main

import (
	"log"
	"net/http"

	"github.com/HunnTeRUS/WebSocketsGoLang/internal/handlers"
	"github.com/HunnTeRUS/WebSocketsGoLang/routes"
)

func main() {
	routes := routes.GetRoutes()

	log.Println("Starting channel listener")

	go handlers.ListenToWsChannel()

	log.Println("Starting webserver on port 8080")

	_ = http.ListenAndServe(":8080", routes)
}

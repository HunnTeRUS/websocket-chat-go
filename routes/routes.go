package routes

import (
	"net/http"

	"github.com/HunnTeRUS/WebSocketsGoLang/internal/handlers"
	"github.com/bmizerany/pat"
)

func GetRoutes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileserver))

	return mux
}

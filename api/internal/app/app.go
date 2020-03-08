package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jschaare/kanji-browser/api/internal/handlers"
)

type App struct {
	Router *mux.Router
}

func (myapp *App) Init() {
	myapp.Router = mux.NewRouter()
	fmt.Println("hi")
	myapp.setRouters()
}

func (myapp *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, myapp.Router))
}

func (myapp *App) setRouters() {
	myapp.Get("/", myapp.requestHandler(handlers.ExampleHandler))
}

func (myapp *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	myapp.Router.HandleFunc(path, f).Methods("GET")
}

type HandlerWrapper func(w http.ResponseWriter, r *http.Request)

func (myapp *App) requestHandler(handler HandlerWrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

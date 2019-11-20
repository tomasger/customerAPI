package customerAPI

import (
	"customerAPI/handlers"
	"customerAPI/storage"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	DB     storage.Database
	Router *mux.Router
}

func NewApp(db storage.Database) *App {
	return &App{
		DB:     db,
		Router: mux.NewRouter(),
	}
}

func (a *App) Init() {
	a.setRouters()
	a.Router.Use(handlers.BasicAuthHandler)
}

func (a *App) Run(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/v1/customers", a.handleRequest(handlers.PostCustomer)).Methods("POST")
	a.Router.HandleFunc("/v1/customers", a.handleRequest(handlers.GetCustomers)).Methods("GET")
	a.Router.HandleFunc("/v1/customers/{id:[0-9]+}", a.handleRequest(handlers.GetCustomer)).Methods("GET")
}

// handleRequest is a wrapper for handlers to provide them with access to Database interface
func (a *App) handleRequest(handler func(storage.Database, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

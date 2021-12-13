package app

import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.initialRoutes()
	return a
}

func (a *App) initialRoutes() {
	api := a.Router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", a.IndexHandler()).Methods("GET")
	api.HandleFunc("/articles", a.GetArticlesHandler()).Methods("GET")
	api.HandleFunc("/article", a.GetArticleHandler()).Methods("GET")
}


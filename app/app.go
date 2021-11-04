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
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/article", a.GetArticleHandler()).Methods("GET")
	a.Router.HandleFunc("/articles", a.GetArticlesHandler()).Methods("GET")
}


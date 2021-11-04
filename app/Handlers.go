package app

import (
	"fmt"
	"log"
	"net/http"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Welcome to the Japan times articles api"); err != nil {
			log.Fatal(err)
		}
	}
}

func (a *App) GetArticleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "This is the handler for getting and scrapping an article"); err != nil {
			log.Fatal(err)
		}
	}
}
func (a *App) GetArticlesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "This is the handler for get all articles for a certain date"); err != nil {
			log.Fatal(err)
		}
	}
}

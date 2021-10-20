package app

import (
	"fmt"
	"log"
	"net/http"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if _, err := fmt.Fprintf(w, "Welcome to the Japan times articles api"); err != nil {
			log.Fatal(err)
		}
	}
}
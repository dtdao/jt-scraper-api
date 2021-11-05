package app

import (
	"encoding/json"
	"fmt"
	"jt-scraper-api/app/models"
	"log"
	"net/http"
	"sync"
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
		var wg sync.WaitGroup
		var article *models.Article
		go func() {
			wg.Add(1)
			article, _ = ScrapeUrl("https://www.japantimes.co.jp/news/2021/11/05/national/ronapreve-covid-drug/")
		}()

		response, _ := json.Marshal(article)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(response)
		if _, err := fmt.Fprintf(w, "This is the handler for get all articles for a certain date"); err != nil {
			log.Fatal(err)
		}
	}
}

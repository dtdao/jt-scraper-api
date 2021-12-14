package app

import (
	"encoding/json"
	"fmt"
	"jt-scraper-api/app/models"
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

func (a *App) GetArticlesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var articles []models.Article
		//if _, err := fmt.Fprintf(w, "This is the handler for getting and scrapping an article"); err != nil {
		//	log.Fatal("error")
		//}
		articles, _ = ScrapeTodaysMainArticles()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(articles)
	}
}
func (a *App) GetArticleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article models.Article
		articleUrl := r.URL.Query().Get("article")
		article, _ = ScrapeUrl(articleUrl)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(article)
	}
}

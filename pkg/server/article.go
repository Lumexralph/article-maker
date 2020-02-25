// Package server has the implementation for the article request handlers.
package server

import (
	"github.com/Lumexralph/article-maker/internal/repository"
	mux "github.com/gorilla/mux"
	"net/http"
)

type ArticleService struct {
	store repository.ArticleRepository
}

// New returns a new HTTP Server for article service
func New(s repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		s,
	}
}

// ServeHTTP helps to implement the ListenAndServe interface
func (as *ArticleService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	as.Router().ServeHTTP(w, r)
}

// Router will handle the routing for the article service
func (as *ArticleService) Router() *mux.Router {
	// TODO: Try to remove trailing slash for request like /article/
	r := mux.NewRouter()

	r.HandleFunc("/article/{id}", as.retrieveArticleHandler).Methods("GET")
	r.HandleFunc("/article/{id}", as.removeAnArticleHandle).Methods("DELETE")
	r.HandleFunc("/article", as.listArticlesHandler).Methods("GET")
	r.HandleFunc("/article", as.updateArticleHandler).Methods("PUT")
	r.HandleFunc("/article", as.createArticleHandler).Methods("POST")

	return r
}

// Package server has the implementation for the
// article request handlers.
package server

import (
	"fmt"
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

func (a *ArticleService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Router().ServeHTTP(w, r)
}

// Router will handle the routing for the article service
func (a *ArticleService) Router() *mux.Router {
	// TODO: Try to remove trailing slash for request like /article/
	r := mux.NewRouter()

	//mux.Handle("/api/", apiHandler{})
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!\n")
	})

	r.HandleFunc("/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("the method >>> ", r.Method, id)

		fmt.Fprintf(w, "Get an article by id %s!\n", id)
	}).Methods("GET")

	r.HandleFunc("/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("the method >>> ", r.Method, id)

		fmt.Fprintf(w, "Delete an article by id %s!\n", id)
	}).Methods("DELETE")

	r.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get all articles!\n")
	}).Methods("GET")

	r.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Update Articles!\n")
	}).Methods("PUT")

	r.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Create new Articles!\n")
	}).Methods("POST")

	return r
}

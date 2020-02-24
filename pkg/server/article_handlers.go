package server

import (
	"encoding/json"
	"fmt"
	"github.com/Lumexralph/article-maker/internal/domain"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func (as *ArticleService) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var a domain.Article

	_ = json.Unmarshal(body, &a)

	// insert the values in the DB
	if err := as.store.CreateArticle(&a); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", body)
}

func (as *ArticleService) updateArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ready to update article\n")
}

func (as *ArticleService) listArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := as.store.ListArticles()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	b, err := json.MarshalIndent(articles, "", "\t")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func retrieveAnArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "Get an article by id %s!\n", id)
}

func removeAnArticleHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("the method >>> ", r.Method, id)

	fmt.Fprintf(w, "Delete an article by id %s!\n", id)
}

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", body)
}

func (as *ArticleService) updateArticleHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var a domain.Article

	_ = json.Unmarshal(body, &a)

	// insert the values in the DB
	if err := as.store.ModifyArticle(&a); err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", body)
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

func (as *ArticleService) retrieveArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	a, err := as.store.GetArticle(id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	b, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (as *ArticleService) removeAnArticleHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := as.store.DeleteArticle(id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Article successfully deleted")
}

// Package repository contains the operations to retrieve or interact with article domain.
// It uses the domain-driven design paradigm such that alternative storage implementations
// may be easily interchanged.
// Reference: https://en.wikipedia.org/wiki/Domain-driven_design
package repository

import (
	"github.com/Lumexralph/article-maker/internal/domain"
)

// ClientRepository interface that any database implementation have to provide
type ArticleRepository interface {
	CreateArticle(*domain.Article) error                               // create an article struct to be passed
	ListArticles() ([]*domain.Article, error)                          // list an article struct to be passed
	GetArticle(string) (*domain.Article, error)                        // get an article by ID
	DeleteArticle(string) error                                        // delete an article using the id
	ModifyArticle(*domain.Article) error                               // update an existing article
	ListArticlesByParameter(...interface{}) ([]*domain.Article, error) // list all articles using parameters
}

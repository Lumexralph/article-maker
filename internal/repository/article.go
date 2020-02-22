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
	//create an article struct to be passed
	CreateArticle(*domain.Article) error
}

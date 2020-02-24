// Package domain contains the entities of the article maker
// business logic
package domain

import "time"

type Article struct {
	ID    int    `json:"-"`
	Title string `json:"title"`
	Body  string `json:"body"` // get just the body field
	Publisher
	Category
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
	Deleted     bool      `json:"-"`
}

type Category struct {
	ID   int    `json:"-"`
	Name string `json:"category"`
}

type Publisher struct {
	ID   int    `json:"-"`
	Name string `json:"publisher"`
}

// Package domain contains the entities of the article maker
// business logic
package domain

import "time"

type Article struct {
	title, body, category, publisher string

	createdAt, publishedAt time.Time
}
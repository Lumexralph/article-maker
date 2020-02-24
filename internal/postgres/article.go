//Package datastore is the implementation of the database session
// in use by the server
package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Lumexralph/article-maker/internal/domain"

	// register driver needed for postgreSQL
	_ "github.com/lib/pq"
)

// CreateClient will create a new database connection with the supplied psqlInfo
func CreateClient(psqlInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// PostgresRepository encapsulates a db connection with the operations
type ArticleStore struct {
	DB    *sql.DB
	Table string
}

// CreateArticle will take the data from the stored file
// and persist it to the database
func (a ArticleStore) CreateArticle(article *domain.Article) error {
	// create category if does not exist
	sqlStatement := `SELECT * FROM category WHERE name=$1;`
	row := a.DB.QueryRow(sqlStatement, article.Category.Name)
	c := new(domain.Category)
	err := row.Scan(&c.ID, &c.Name)
	if err == sql.ErrNoRows {
		sqlStatement := `INSERT INTO category (name) 
						VALUES ($1);`
		if _, err := a.DB.Exec(sqlStatement, article.Category.Name); err != nil {
			return err
		}
	}

	// create publisher if it does not exist
	sqlStatement = `SELECT * FROM publisher WHERE name=$1;`
	row = a.DB.QueryRow(sqlStatement, article.Publisher.Name)
	p := new(domain.Publisher)
	err = row.Scan(&p.ID, &p.Name)
	if err == sql.ErrNoRows {
		sqlStatement := `INSERT INTO publisher (name) 
						VALUES 
						($1);`
		if _, err := a.DB.Exec(sqlStatement, article.Publisher.Name); err != nil {
			return err
		}
	}
	fmt.Println(c.Name, p.Name)
	// create an article
	sqlStatement = `INSERT INTO article (title, body, category, publisher, created_at, published_at, deleted)
 					VALUES 
 					($1, $2, $3, $4, $5, $6, $7);`
	if _, err := a.DB.Exec(
		sqlStatement,
		article.Title,
		article.Body,
		article.Category.Name,
		article.Publisher.Name,
		article.CreatedAt,
		article.PublishedAt,
		false,
	); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (a ArticleStore) ListArticles() ([]*domain.Article, error) {
	rows, err := a.DB.Query("SELECT title, body, category, publisher, created_at, published_at, deleted FROM article")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]*domain.Article, 0)
	for rows.Next() {
		a := new(domain.Article)
		err := rows.Scan(&a.Title, &a.Body, &a.Category.Name, &a.Publisher.Name, &a.CreatedAt, &a.PublishedAt, &a.Deleted)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

func (a ArticleStore) GetArticle(id string) (*domain.Article, error) {
	sqlStatement := `SELECT title, body, category, publisher, created_at, published_at, deleted FROM article WHERE id=$1;`
	row := a.DB.QueryRow(sqlStatement, id)
	da := new(domain.Article)
	err := row.Scan(&da.Title, &da.Body, &da.Category.Name, &da.Publisher.Name, &da.CreatedAt, &da.PublishedAt, &da.Deleted)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return da, nil
}

func (a ArticleStore) DeleteArticle(id string) error {
	sqlStatement := `DELETE FROM article WHERE id=$1;`
	_, err := a.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}
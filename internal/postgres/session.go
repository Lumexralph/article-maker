//Package datastore is the implementation of the database session
// in use by the server
package postgres

import (
	"database/sql"
	"github.com/Lumexralph/article-maker/internal/domain"
)

// CreateClient will create a new database connection with the supplied psqlInfo
func CreateClient(psqlInfo string) (*sql.DB, error){
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db, nil
}

// DB encapsulates a db connection with the operations
type PostgresRepository struct {
	db *sql.DB
	table string
}

// CreateFileMetaData will take the data from the stored file
// and persist it to the database
func (pg *PostgresRepository) CreateArticle(article *domain.Article) {
	sqlStatement := `INSERT INTO ` + pg.table  + ` (name, slug, format, path, size)
	VALUES ($1, $2, $3, $4, $5)`
	if _, err := pg.db.Exec(sqlStatement, article., fd.slug, fd.format, fd.path, fd.size); err != nil {
		panic(err)
	}
}
package server

import (
	"github.com/Lumexralph/article-maker/internal/domain"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDB struct{}

func (m mockDB) ListArticles() ([]*domain.Article, error) {
	articles := []*domain.Article{
		{
			ID:        1,
			Title:     "game of soccer",
			Body:      "Soccer is a game",
			Category:  domain.Category{Name: "sport"},
			Publisher: domain.Publisher{Name: "John"},
		},
		{
			ID:        2,
			Title:     "game of life",
			Body:      "Life is a game",
			Category:  domain.Category{Name: "life"},
			Publisher: domain.Publisher{Name: "Femi"},
		},
	}

	return articles, nil
}

func (m mockDB) CreateArticle(article *domain.Article) error {
	return nil
}

func (m mockDB) ModifyArticle(article *domain.Article) error {
	return nil
}

func (m mockDB) ListArticlesByParameter(fields ...interface{}) ([]*domain.Article, error) {
	articles := []*domain.Article{
		{
			ID:        1,
			Title:     "game of soccer",
			Body:      "Soccer is a game",
			Category:  domain.Category{Name: "sport"},
			Publisher: domain.Publisher{Name: "John"},
		},
	}

	return articles, nil
}

func (m mockDB) GetArticle(id string) (*domain.Article, error) {

	return &domain.Article{
		ID:        1,
		Title:     "game of soccer",
		Body:      "Soccer is a game",
		Category:  domain.Category{Name: "sport"},
		Publisher: domain.Publisher{Name: "John"},
	}, nil
}

func (m mockDB) DeleteArticle(id string) error {
	return nil
}

var serv = New(mockDB{})

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestListArticlesHandler(t *testing.T) {
	t.Run("GET /article", func(t *testing.T) {
		r, err := http.NewRequest("GET", "/article", nil)
		checkError(err, t)

		want := `[
	{
		"ID": 1,
		"title": "game of soccer",
		"body": "Soccer is a game",
		"publisher": "John",
		"category": "sport",
		"created_at": "0001-01-01T00:00:00Z",
		"published_at": "0001-01-01T00:00:00Z"
	},
	{
		"ID": 2,
		"title": "game of life",
		"body": "Life is a game",
		"publisher": "Femi",
		"category": "life",
		"created_at": "0001-01-01T00:00:00Z",
		"published_at": "0001-01-01T00:00:00Z"
	}
]`

		rr := httptest.NewRecorder()
		http.HandlerFunc(serv.listArticlesHandler).ServeHTTP(rr, r)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("listArticlesHandler() GET /article got wrong status code. got %d; want %d\n", http.StatusOK, status)
		}

		if diff := cmp.Diff(want, rr.Body.String()); diff != "" {
			t.Errorf("listArticlesHandler() GET /article, mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("GET /article with query parameters", func(t *testing.T) {
		// request with query parameters
		r, err := http.NewRequest("GET", "/article?category=Commercials&publisher=John", nil)
		checkError(err, t)

		rr := httptest.NewRecorder()
		http.HandlerFunc(serv.listArticlesHandler).ServeHTTP(rr, r)

		want := `[
	{
		"ID": 1,
		"title": "game of soccer",
		"body": "Soccer is a game",
		"publisher": "John",
		"category": "sport",
		"created_at": "0001-01-01T00:00:00Z",
		"published_at": "0001-01-01T00:00:00Z"
	}
]`

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("listArticlesHandler() GET /article?category=Commercials&publisher=John got wrong status code. got %d; want %d\n", http.StatusOK, status)
		}

		if diff := cmp.Diff(want, rr.Body.String()); diff != "" {
			t.Errorf("listArticlesHandler() mismatch GET /article?category=Commercials&publisher=John, (-want +got):\n%s", diff)
		}
	})
}

func TestRetrieveArticleHandler(t *testing.T) {
	r, err := http.NewRequest("GET", "/article/1", nil)
	checkError(err, t)

	want := `{
	"ID": 1,
	"title": "game of soccer",
	"body": "Soccer is a game",
	"publisher": "John",
	"category": "sport",
	"created_at": "0001-01-01T00:00:00Z",
	"published_at": "0001-01-01T00:00:00Z"
}`

	rr := httptest.NewRecorder()

	http.HandlerFunc(serv.retrieveArticleHandler).ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("retrieveArticleHandler(1) GET /article/1 got wrong status code. got %d; want %d\n", http.StatusOK, status)
	}

	if diff := cmp.Diff(want, rr.Body.String()); diff != "" {
		t.Errorf("retrieveArticleHandler(1) GET /article/1, mismatch (-want +got):\n%s", diff)
	}
}

func TestRemoveAnArticleHandler(t *testing.T) {
	r, err := http.NewRequest("DELETE", "/article/1", nil)
	checkError(err, t)

	rr := httptest.NewRecorder()

	want := "Article successfully deleted"

	http.HandlerFunc(serv.removeAnArticleHandler).ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("removeAnArticleHandler(1) DELETE /article/1 got wrong status code. got %d; want %d\n", http.StatusOK, status)
	}

	if diff := cmp.Diff(want, rr.Body.String()); diff != "" {
		t.Errorf("retrieveArticleHandler(1) DELETE /article/1, mismatch (-want +got):\n%s", diff)
	}
}

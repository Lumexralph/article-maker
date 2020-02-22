package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFetchURL(t *testing.T) {
	// a mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "list of comments from the server")
	}))
	defer ts.Close()

	t.Run("valid url", func(t *testing.T) {
		got, err := FetchURL(ts.URL)
		if err != nil {
			t.Errorf("FetchURL(%q) got=%v; want=%v", ts.URL, err, nil)
		}

		want := "list of comments from the server"
		if string(got) == want {
			t.Errorf("FetchURL(%q) got=%v; want=%q", ts.URL, got, want)
		}
	})

	t.Run("invalid url", func(t *testing.T) {
		got, err := FetchURL("article-maker")
		if err == nil {
			t.Errorf("FetchURL(%q) should return error, got=%v;", ts.URL, nil)
		}

		if got != nil {
			t.Errorf("FetchURL(%q) got=%v; want=%v", "article-maker", got, nil)
		}
	})
}

func TestParseComment(t *testing.T) {
	cases := []struct {
		name         string
		responseBody []byte
		want         []Comment
	}{
		{
			name:         "valid parsing json comment from request body",
			responseBody: []byte(`[{"body": "list of comments from the server"}]`),
			want: []Comment{
				{Body: "list of comments from the server"},
			},
		},
		{
			name:         "invalid parsing json comment from response body",
			responseBody: []byte(`[{"email": "list of comments from the server"}]`),
			want:         []Comment{{}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseComment(tc.responseBody)
			if err != nil {
				t.Errorf("ParseComment(%v) got=%v; want=%v", tc.want, err, nil)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("ParseComment() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGenerateWordCount(t *testing.T) {
	comments := []Comment{
		{Body: "list of comments from of server"},
	}
	want := map[string]int{
		"list":     1,
		"of":       2,
		"comments": 1,
		"from":     1,
		"server":   1,
	}

	got := GenerateWordCount(comments)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GenerateWordCount() mismatch (-want +got):\n%s", diff)
	}
}

func TestLeastCommonWords(t *testing.T) {
	records := map[string]int{
		"list":     1,
		"comments": 2,
		"book":     4,
		"game":     5,
		"films":    3,
	}
	want := []record{
		{Word: "list", Count: 1},
		{Word: "comments", Count: 2},
		{Word: "films", Count: 3},
		{Word: "book", Count: 4},
	}

	got := LeastCommonWords(records, 4)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("LeastCommonWords() mismatch (-want +got):\n%s", diff)
	}
}

func TestWordCountToJSON(t *testing.T) {
	records := []record{
		{Word: "list", Count: 1},
		{Word: "comments", Count: 2},
	}

	want, _ := json.MarshalIndent(records, "", "\t")
	got, err := WordCountToJSON(records)
	if err != nil {
		t.Errorf("WordCountToJSON(%v) got=%v; want=%v", records, err, nil)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("WordCountToJSON() mismatch (-want +got):\n%s", diff)
	}
}

/*
Copyright © 2020 OLUMIDE OGUNDELE <olumideralph@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Package pkg holds the reusable functions, types and methods
used by the cmd package implementations.
*/
package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// FetchURL parses the supplied URL and makes a GET request to the URL
func FetchURL(rawURL string) ([]byte, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, nil
}

// Comment to encode the response from json
type Comment struct {
	Body string `json:"body"` // get just the body field
}

// ParseComment encodes the response to a slice of struct
func ParseComment(body []byte) ([]Comment ,error) {
	var c []Comment

	err := json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// GenerateWordCount get a generated count of all the words in the body field
func GenerateWordCount(comments []Comment) map[string]int {
	records := make(map[string]int)

	for _, comment := range comments {
		separatedWords := strings.Fields(comment.Body)
		for _, word := range separatedWords {
			records[word] += 1
		}
	}
	return records
}

type record struct {
	Word string `json:"word"`
	Count int `json:"count"`
}

// to be able to sort the records
type wordRecord struct {
	records []record
}

// Len is part of sort.Interface.
func (wr wordRecord) Len() int {
	return len(wr.records)
}

// Swap is part of sort.Interface.
func (wr wordRecord) Swap(i, j int) {
	wr.records[i], wr.records[j] = wr.records[j], wr.records[i]
}

// Less is part of sort.Interface.
func (wr wordRecord) Less(i, j int) bool {
	return wr.records[i].Count < wr.records[j].Count
}

// LeastCommonWords will get the number of least word occurrence with n as the
// required number of words to get
func LeastCommonWords(records map[string]int, n int) []record {
	wr := wordRecord{}

	for word, count := range records {
		wr.records = append(wr.records, record{Word: word, Count: count})
	}

	// sort the records according to count
	sort.Sort(wr)

	return wr.records[:n]
}

// WordCountToJSON decodes the word count to JSON
func WordCountToJSON(records []record) ([]byte, error) {
	b, err := json.MarshalIndent(records, "", "\t")
	if err != nil {
		return nil, err
	}

	return b, nil
}

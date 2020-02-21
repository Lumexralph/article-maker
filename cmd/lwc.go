/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

var urlFlag string

func init() {
	rootCmd.AddCommand(lwcCmd)

	// Here you will define your flags and configuration settings.
	lwcCmd.Flags().StringVarP(&urlFlag, "url", "u", "", "URL to get the words or string from")
}

// lwcCmd represents the lwc command
var lwcCmd = &cobra.Command{
	Use:   "lwc",
	Short: "find the four least used words and their word count",
	Long: `fetches the comments data from the supplied url. In the ‘’body’’ field, 
			finds the four least used words and their word count.`,
	Run: leastWordCount,
}

func leastWordCount(cmd *cobra.Command, args []string) {
	resp, err := fetchURL(urlFlag)
	if err != nil {
		fmt.Println(err)
		return
	}

	// return error or the computed value from the request body using goroutines
	fmt.Println(string(resp))
}

// fetchURL parses the supplied URL and makes a GET request to the URL
func fetchURL(rawURL string) ([]byte, error) {
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

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
*/
package cmd

import (
	"fmt"

	"github.com/Lumexralph/article-maker/pkg"

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
	resp, err := pkg.FetchURL(urlFlag)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := pkg.ParseComment(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	// generate the word counts
	records := pkg.GenerateWordCount(c)
	// get the least common words
	wordRecords := pkg.LeastCommonWords(records, 4)
	b, err := pkg.WordCountToJSON(wordRecords)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

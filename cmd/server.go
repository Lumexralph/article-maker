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
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var portFlag string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start the server to expose the APIs",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Run: server,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.
	serverCmd.Flags().StringVarP(
		&portFlag,
		"port",
		"p",
		"",
		"port to start the server on",
	)

}

func server(cmd *cobra.Command, args []string) {
	fmt.Println(portFlag)
	//port := os.Getenv("PORT")
	log.Printf("Starting server on port: %s \n", portFlag)
	err := http.ListenAndServe(":"+portFlag, nil)
	if err != nil {
		fmt.Println(err)
	}
}

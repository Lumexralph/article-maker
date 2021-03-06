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
	"github.com/Lumexralph/article-maker/internal/postgres"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"net/http"
	"os"

	articleserver "github.com/Lumexralph/article-maker/pkg/server"
	log "github.com/golang/glog"
)

var portFlag string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start the server to expose the APIs",
	Long: `server sub-command starts a server on the provided port to
			listen from and supplies endpoints to perform different
			operations with articles.`,
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
	err := godotenv.Load()
	if err != nil {
		log.Info("Could not load env. file...")
	}

	// start the service
	log.Infof("Starting server on port: %s \n", portFlag)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// create database url
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
	)

	db, err := postgres.CreateClient(connStr)
	if err != nil {
		log.Info(err)
	}
	defer db.Close()

	// article service datastore
	articleStore := postgres.ArticleStore{
		DB: db,
	}

	// create new article service
	serv := articleserver.New(articleStore)

	err = http.ListenAndServe(":"+portFlag, serv)
	if err != nil {
		log.Info(err)
	}

}

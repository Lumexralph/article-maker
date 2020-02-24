# Article Maker

An implementation of an article maker system with a CLI tool to fetch comments with a supplied URL and an API for articles.

## Getting started

Clone the repository

    git clone https://github.com/Lumexralph/article-maker.git
    
Change into the directory

    cd uconv

Install all the necessary packages, run

    go mod tidy

Build and install the tool in your computer

    go install .

article-maker gets added as a binary in your computer and you can start using it.

It gives you access to 2 tools;

1. lwc - lowest common words in a list of comments

        article-maker lwc --url https://jsonplaceholder.typicode.com/comments
    
    a json reponse gets printed to the stdout or the terminal like below
    
        [
                {
                        "word": "amet",
                        "count": 24
                },
                {
                        "word": "ducimus",
                        "count": 31
                },
                {
                        "word": "dolore",
                        "count": 35
                },
                {
                        "word": "assumenda",
                        "count": 35
                }
        ]
2. server - this handles the article maker api.
    Start the server using the following command as below
    
        article-maker server --port 5000  // you can supply any port of choice
        
     
  Then you have access to the following endpoints or URL
  
Feature | Method | URL | Payload
-------- | ------- | ------- | --------
Create an article | POST | /article | { "title": "Game of thrones", "body": "Lorem ipsum dolor sit amet, "category": "Commercials", "publisher": "John Wale", "created_at": "2006-01-02T15:04:05+07:00", "published_at": "2006-01-02T15:04:05+07:00" }
Get all articles | GET | /article
Update an article | PUT | /article {"id" : 5, ... POST payload}
Get a specific article by id | GET | /article/:id
Delete a specific article | DELETE | /article/:id


For the api, you'll need a postgres database.
## Errors and bugs

If something is not behaving intuitively, it is a bug and should be reported.
Report it here by creating an issue: https://github.com/Lumexralph/article-maker/issues

Copyright (c) 2020 LumexRalph. Released under the [Apache License](https://github.com/Lumexralph/article-maker/blob/master/LICENSE).


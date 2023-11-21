package main

import (
	"os"

	"github.com/lastvoidtemplar/github-api-client/infra"
	"github.com/lastvoidtemplar/github-api-client/present"
)

func main() {
	usernames := infra.ReadUsernamesFromFile(os.Args[1])
	users := infra.FetchUsersFromGithubApi(usernames)

	tableContent := present.CreateTableContent(users)
	present.PrintTable(tableContent)
}

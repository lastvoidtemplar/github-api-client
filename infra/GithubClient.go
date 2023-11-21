package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/lastvoidtemplar/github-api-client/domain"
)

func FetchUsersFromGithubApi(usernames []string) []domain.User {
	users := make([]domain.User, len(usernames))
	sync := make(chan struct{})
	for ind := 0; ind < len(usernames); ind++ {
		go fetchUserFromGithubApi(usernames[ind], &users[ind], sync)
	}
	for ind := 0; ind < len(usernames); ind++ {
		<-sync
	}
	return users
}

func fetchUserFromGithubApi(username string, user *domain.User, sync chan struct{}) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s", username), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "Accept: application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Print(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	fetchUserReposFromGithubApi(user)
	fetchReposLangFromGithubApi(user.Repos)
	sync <- struct{}{}
}

func fetchUserReposFromGithubApi(user *domain.User) {
	req, err := http.NewRequest("GET", fmt.Sprintf(user.ReposURL), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "Accept: application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Print(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &user.Repos); err != nil {
		log.Fatal(err)
	}
}

func fetchReposLangFromGithubApi(repos []domain.Repo) {

	//execute request at the same time
	sync := make(chan struct{})
	for ind := 0; ind < len(repos); ind++ {
		go func(sync chan struct{}, ind int) {
			req, err := http.NewRequest("GET", repos[ind].LanguagesURL, nil)
			if err != nil {
				log.Fatal(err)
			}
			req.Header.Add("Accept", "Accept: application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			if resp.StatusCode != 200 {
				log.Print(resp.Status)
			}

			body, err := io.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			if err := json.Unmarshal(body, &repos[ind].Langs); err != nil {
				log.Fatal(err)
			}
			sync <- struct{}{}
		}(sync, ind)
	}
	//wait all request to finish
	for ind := 0; ind < len(repos); ind++ {
		<-sync
	}
}

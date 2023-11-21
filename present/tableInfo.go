package present

import (
	"encoding/json"
	"math"
	"time"

	"github.com/lastvoidtemplar/github-api-client/domain"
	"github.com/lastvoidtemplar/github-api-client/utils"
)

type LangsDisto map[string]float64

type YearActivityDistro map[int]int

type TableInfo struct {
	Login              string
	Name               string
	SiteAdmin          bool
	PublicRepos        int
	Forks              int
	LangsDistro        string
	Followers          int
	Following          int
	YearActivityDistro string
	CreatedAt          time.Time
}

func createLangsDistro(user *domain.User) string {
	langsDistro := make(LangsDisto)
	var count int64
	for ind := 0; ind < len(user.Repos); ind++ {
		for k, v := range user.Repos[ind].Langs {
			langsDistro[k] += float64(v)
			count += v
		}
	}
	for k, v := range langsDistro {
		langsDistro[k] = math.Round(v/float64(count)*100*100) / 100
	}
	jsonRes, _ := json.Marshal(langsDistro)
	return string(jsonRes)
}

func createYearActiveDistro(user *domain.User) string {
	yearActivityDistro := make(YearActivityDistro)

	for ind := 0; ind < len(user.Repos); ind++ {
		for year := user.Repos[ind].CreatedAt.Year(); year <= user.Repos[ind].UpdatedAt.Year(); year++ {
			yearActivityDistro[year]++
		}
	}
	jsonRes, _ := json.Marshal(yearActivityDistro)
	return string(jsonRes)
}

func CreateTableContent(users []domain.User) []TableInfo {
	tableContent := make([]TableInfo, 0, len(users))

	for ind := 0; ind < len(users); ind++ {

		tableInfo := TableInfo{
			Login:       users[ind].Login,
			Name:        users[ind].Name,
			SiteAdmin:   users[ind].SiteAdmin,
			PublicRepos: users[ind].PublicRepos,
			Forks: len(utils.Filter(users[ind].Repos, func(el *domain.Repo) bool {
				return el.Fork
			})),
			LangsDistro:        createLangsDistro(&users[ind]),
			Followers:          users[ind].Followers,
			Following:          users[ind].Following,
			YearActivityDistro: createYearActiveDistro(&users[ind]),
			CreatedAt:          users[ind].CreatedAt,
		}

		tableContent = append(tableContent, tableInfo)
	}
	return tableContent
}

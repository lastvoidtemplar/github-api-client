package domain

import "time"

type Repo struct {
	ID              int    `json:"id"`
	NodeID          string `json:"node_id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Private         bool   `json:"private"`
	HTMLURL         string `json:"html_url"`
	Description     any    `json:"description"`
	Fork            bool   `json:"fork"`
	Langs           Langs
	LanguagesURL    string    `json:"languages_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
	GitURL          string    `json:"git_url"`
	CloneURL        string    `json:"clone_url"`
	SvnURL          string    `json:"svn_url"`
	Homepage        any       `json:"homepage"`
	Size            int       `json:"size"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Language        string    `json:"language"`
	HasIssues       bool      `json:"has_issues"`
	HasProjects     bool      `json:"has_projects"`
	HasDownloads    bool      `json:"has_downloads"`
	HasWiki         bool      `json:"has_wiki"`
	HasPages        bool      `json:"has_pages"`
	HasDiscussions  bool      `json:"has_discussions"`
	ForksCount      int       `json:"forks_count"`
	Archived        bool      `json:"archived"`
	Disabled        bool      `json:"disabled"`
	OpenIssuesCount int       `json:"open_issues_count"`
	License         any       `json:"license"`
	AllowForking    bool      `json:"allow_forking"`
	IsTemplate      bool      `json:"is_template"`
	Visibility      string    `json:"visibility"`
	Forks           int       `json:"forks"`
	OpenIssues      int       `json:"open_issues"`
	Watchers        int       `json:"watchers"`
	DefaultBranch   string    `json:"default_branch"`
}

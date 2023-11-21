package domain

import "time"

type User struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	AvatarURL   string `json:"avatar_url"`
	HTMLURL     string `json:"html_url"`
	Repos       []Repo
	ReposURL    string    `json:"repos_url"`
	Type        string    `json:"type"`
	SiteAdmin   bool      `json:"site_admin"`
	Name        string    `json:"name"`
	Company     any       `json:"company"`
	Blog        string    `json:"blog"`
	Location    any       `json:"location"`
	Email       any       `json:"email"`
	Hireable    any       `json:"hireable"`
	Bio         any       `json:"bio"`
	PublicRepos int       `json:"public_repos"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	CreatedAt   time.Time `json:"created_at"`
}

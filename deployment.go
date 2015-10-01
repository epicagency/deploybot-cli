package main

type Deployment struct {
	Id                   int
	RepositoryId         int    `json:"repository_id"`
	EnvironmentId        int    `json:"environment_id"`
	UserId               int    `json:"user_id"`
	DeployedVersion      string `json:"deployed_version"`
	DeployFromScratch    bool   `json:"deploy_from_scratch"`
	TriggerNotifications bool   `json:"trigger_notifications"`
	IsAutomatic          bool   `json:"is_automatic"`
	Comment              string
	AuthorName           string `json:"author_name"`
	State                string
	Retries              int
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}

type Deployments struct {
	Meta    Meta
	Entries []Deployment
}

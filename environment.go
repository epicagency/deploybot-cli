package main

type Environment struct {
	Id                int
	Name              string
	CurrentVersion    string `json:"current_version"`
	IsAutomatic       bool   `json:"is_automatic"`
	RepositoryId      int    `json:"repository_id"`
	BranchName        string `json:"branch_name"`
	DeployWebhookUrl  string `json:"deploy_webhook_url"`
	StatusBadgePngUrl string `json:"status_badge_png_url"`
	StatusBadgeSvgUrl string `json:"status_badge_svg_url"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type Environments struct {
	Meta    Meta
	Entries []Environment
}

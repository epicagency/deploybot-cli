package main

type Repository struct {
	Id                int
	Name              string
	Title             string
	ColorLabel        string `json:"color_label"`
	Type              string
	RefreshWebhookUrl string `json:"refresh_webhook_url"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type Repositories struct {
	Meta    Meta
	Entries []Repository
}

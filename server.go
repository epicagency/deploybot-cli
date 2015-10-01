package main

type Server struct {
	Id             int
	Name           string
	Protocol       string
	RepositoryId   int    `json:"repository_id"`
	EnvironmentId  int    `json:"environment_id"`
	PreDeployHook  string `json:"pre_deploy_hook"`
	PostDeployHook string `json:"post_deploy_hook"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type Servers struct {
	Meta    Meta
	Entries []Server
}

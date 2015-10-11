package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Meta struct {
	Next    int
	NextUrl string `json:"next_url"`
	Total   int
}

type DeployBot struct {
	Config *Config
}

func (d *DeployBot) getRawContent(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (d *DeployBot) getContent(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s.deploybot.com/api/v1/%s", d.Config.Domain, url), nil)
	req.Header.Add("X-Api-Token", d.Config.Token)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (d *DeployBot) postContent(url string, post_body string) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s.deploybot.com/api/v1/%s", d.Config.Domain, url), ioutil.NopCloser(strings.NewReader(post_body)))
	req.Header.Add("X-Api-Token", d.Config.Token)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (d *DeployBot) fetch(path string, v interface{}) error {
	content, err := d.getContent(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, v)
}

func (d *DeployBot) GetRepositories() (*Repositories, error) {
	repositories := new(Repositories)
	err := d.fetch("repositories", repositories)
	if err != nil {
		return nil, err
	}
	return repositories, err
}

func (d *DeployBot) GetRepository(id int) (*Repository, error) {
	repository := new(Repository)
	err := d.fetch(fmt.Sprintf("repositories/%d", id), repository)
	if err != nil {
		return nil, err
	}
	return repository, err
}

func (d *DeployBot) GetEnvironments() (*Environments, error) {
	environments := new(Environments)
	err := d.fetch("environments", environments)
	if err != nil {
		return nil, err
	}
	return environments, err
}

func (d *DeployBot) GetEnvironmentsByRepository(repositoryId int) (*Environments, error) {
	environments := new(Environments)
	err := d.fetch(fmt.Sprintf("environments?repository_id=%d", repositoryId), environments)
	if err != nil {
		return nil, err
	}
	return environments, err
}

func (d *DeployBot) GetEnvironment(id int) (*Environment, error) {
	environment := new(Environment)
	err := d.fetch(fmt.Sprintf("environments/%d", id), environment)
	if err != nil {
		return nil, err
	}
	return environment, err
}

func (d *DeployBot) GetDeploymentsByRepository(repositoryId int) (*Deployments, error) {
	deployments := new(Deployments)
	err := d.fetch(fmt.Sprintf("repository_id=%d", repositoryId), deployments)
	if err != nil {
		return nil, err
	}
	return deployments, err
}

func (d *DeployBot) GetDeploymentsByEnvironment(environmentId int) (*Deployments, error) {
	deployments := new(Deployments)
	err := d.fetch(fmt.Sprintf("deployments?environment_id=%d", environmentId), deployments)
	if err != nil {
		return nil, err
	}
	return deployments, err
}

func (d *DeployBot) GetDeployment(id int) (*Deployment, error) {
	deployment := new(Deployment)
	err := d.fetch(fmt.Sprintf("deployments/%d", id), deployment)
	if err != nil {
		return nil, err
	}
	return deployment, err
}

func (d *DeployBot) TriggerDeployment(settings DeploymentSetting) (*Deployment, error) {
	body, _ := json.Marshal(settings)
	content, err := d.postContent("deployments", string(body))
	if err != nil {
		return nil, err
	}
	var deployment Deployment
	err = json.Unmarshal(content, &deployment)
	if err != nil {
		return nil, err
	}

	return &deployment, err
}

func (d *DeployBot) GetUsers() (*Users, error) {
	users := new(Users)
	err := d.fetch("users", users)
	if err != nil {
		return nil, err
	}
	return users, err
}

func (d *DeployBot) GetUser(id int) (*User, error) {
	user := new(User)
	err := d.fetch(fmt.Sprintf("users/%d", id), user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (d *DeployBot) GetServers() (*Servers, error) {
	servers := new(Servers)
	err := d.fetch("servers", servers)
	if err != nil {
		return nil, err
	}
	return servers, err
}

func (d *DeployBot) GetServersByEnvironment(environmentId int) (*Servers, error) {
	servers := new(Servers)
	err := d.fetch(fmt.Sprintf("servers?environment_id=%d", environmentId), servers)
	if err != nil {
		return nil, err
	}
	return servers, err
}

func (d *DeployBot) GetServersByRepository(repositoryId int) (*Servers, error) {
	servers := new(Servers)
	err := d.fetch(fmt.Sprintf("servers?repository_id=%d", repositoryId), servers)
	if err != nil {
		return nil, err
	}
	return servers, err
}

func (d *DeployBot) GetServer(id int) (*Server, error) {
	server := new(Server)
	err := d.fetch(fmt.Sprintf("servers/%d", id), server)
	if err != nil {
		return nil, err
	}
	return server, err
}

func (d *DeployBot) Refresh(repository *Repository) error {
	content, err := d.getRawContent(repository.RefreshWebhookUrl)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}

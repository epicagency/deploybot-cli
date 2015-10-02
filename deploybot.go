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

func (d *DeployBot) GetRepositories() (*Repositories, error) {
	content, err := d.getContent("repositories")
	if err != nil {
		return nil, err
	}
	var repositories Repositories
	err = json.Unmarshal(content, &repositories)
	if err != nil {
		return nil, err
	}

	return &repositories, err
}

func (d *DeployBot) GetRepository(id int) (*Repository, error) {
	content, err := d.getContent(fmt.Sprintf("repositories/%d", id))
	if err != nil {
		return nil, err
	}
	var repository Repository
	err = json.Unmarshal(content, &repository)
	if err != nil {
		return nil, err
	}

	return &repository, err
}

func (d *DeployBot) GetEnvironments() (*Environments, error) {
	content, err := d.getContent("environments")
	if err != nil {
		return nil, err
	}
	var environments Environments
	err = json.Unmarshal(content, &environments)
	if err != nil {
		return nil, err
	}

	return &environments, err
}

func (d *DeployBot) GetEnvironmentsByRepository(repositoryId int) (*Environments, error) {
	content, err := d.getContent(fmt.Sprintf("environments?repository_id=%d", repositoryId))
	if err != nil {
		return nil, err
	}
	var environments Environments
	err = json.Unmarshal(content, &environments)
	if err != nil {
		return nil, err
	}

	return &environments, err
}

func (d *DeployBot) GetEnvironment(id int) (*Environment, error) {
	content, err := d.getContent(fmt.Sprintf("environments/%d", id))
	if err != nil {
		return nil, err
	}
	var environment Environment
	err = json.Unmarshal(content, &environment)
	if err != nil {
		return nil, err
	}

	return &environment, err
}

func (d *DeployBot) GetDeploymentsByRepository(repositoryId int) (*Deployments, error) {
	content, err := d.getContent(fmt.Sprintf("repository_id=%d", repositoryId))
	if err != nil {
		return nil, err
	}
	var deployments Deployments
	err = json.Unmarshal(content, &deployments)
	if err != nil {
		return nil, err
	}

	return &deployments, err
}

func (d *DeployBot) GetDeploymentsByEnvironment(environmentId int) (*Deployments, error) {
	content, err := d.getContent(fmt.Sprintf("deployments?environment_id=%d", environmentId))
	if err != nil {
		return nil, err
	}
	var deployments Deployments
	err = json.Unmarshal(content, &deployments)
	if err != nil {
		return nil, err
	}

	return &deployments, err
}

func (d *DeployBot) GetDeployment(id int) (*Deployment, error) {
	content, err := d.getContent(fmt.Sprintf("deployments/%d", id))
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
	content, err := d.getContent("users")
	if err != nil {
		return nil, err
	}
	var users Users
	err = json.Unmarshal(content, &users)
	if err != nil {
		return nil, err
	}

	return &users, err
}

func (d *DeployBot) GetUser(id int) (*User, error) {
	content, err := d.getContent(fmt.Sprintf("users/%d", id))
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(content, &user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (d *DeployBot) GetServers() (*Servers, error) {
	content, err := d.getContent("servers")
	if err != nil {
		return nil, err
	}
	var servers Servers
	err = json.Unmarshal(content, &servers)
	if err != nil {
		return nil, err
	}

	return &servers, err
}

func (d *DeployBot) GetServersByEnvironment(environmentId int) (*Servers, error) {
	content, err := d.getContent(fmt.Sprintf("servers?environment_id=%d", environmentId))
	if err != nil {
		return nil, err
	}
	var servers Servers
	err = json.Unmarshal(content, &servers)
	if err != nil {
		return nil, err
	}

	return &servers, err
}

func (d *DeployBot) GetServersByRepository(repositoryId int) (*Servers, error) {
	content, err := d.getContent(fmt.Sprintf("servers?repository_id=%d", repositoryId))
	if err != nil {
		return nil, err
	}
	var servers Servers
	err = json.Unmarshal(content, &servers)
	if err != nil {
		return nil, err
	}

	return &servers, err
}

func (d *DeployBot) GetServer(id int) (*Server, error) {
	content, err := d.getContent(fmt.Sprintf("servers/%d", id))
	if err != nil {
		return nil, err
	}
	var server Server
	err = json.Unmarshal(content, &server)
	if err != nil {
		return nil, err
	}

	return &server, err
}

func (d *DeployBot) Refresh(repository *Repository) error {
	content, err := d.getRawContent(repository.RefreshWebhookUrl)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}

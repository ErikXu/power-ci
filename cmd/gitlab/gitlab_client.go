package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"power-ci/models/gitlab"
)

type GitlabClient struct {
	Client      http.Client
	Host        string
	AccessToken string
}

func (client *GitlabClient) GrantOauthToken(username string, password string) gitlab.OauthResponse {
	request := &gitlab.OauthRequest{
		GrantType: "password",
		Username:  username,
		Password:  password,
	}

	return call(client, "POST", client.Host, "/oauth/token", *request)
}

func (client *GitlabClient) GetUserByUsername(username string) []gitlab.GetUserResponse {
	url := fmt.Sprintf("%s/api/v4/users?username=%s", client.Host, username)
	res, err := client.Client.Get(url)
	if err != nil {
		panic("Call rest api failed")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	response := []gitlab.GetUserResponse{}
	json.Unmarshal(body, &response)
	return response
}

func (client *GitlabClient) CreateUser(admin bool, username string, name string, email string, password string) gitlab.CreateUserResponse {
	url := fmt.Sprintf("%s/api/v4/users", client.Host)

	request := &gitlab.CreateUserRequest{
		Admin:    admin,
		Username: username,
		Name:     name,
		Email:    email,
		Password: password,
	}

	jsonValue, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Authorization", "Bearer "+client.AccessToken)

	res, err := client.Client.Do(req)
	if err != nil {
		panic("Call rest api failed")
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	response := gitlab.CreateUserResponse{}
	json.Unmarshal(body, &response)
	return response
}

func call[Request gitlab.GitlabRequest, Response gitlab.GitlabResponse](client *GitlabClient, method string, host string, api string, request Request) Response {
	url := host + api

	jsonValue, _ := json.Marshal(request)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	res, err := client.Client.Do(req)
	if err != nil {
		panic("Call rest api failed")
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	response := Response{}
	json.Unmarshal(body, &response)

	return response
}

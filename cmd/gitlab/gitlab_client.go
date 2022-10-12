package gitlab

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"power-ci/models/gitlab"
	"strings"
)

type GitlabClient struct {
	Client      http.Client
	AccessToken string
}

func (client *GitlabClient) GrantOauthToken(method string, host string, request gitlab.OauthRequest) gitlab.OauthResponse {
	return call(client, method, host, "/oauth/token", request)
}

func call[Request gitlab.GitlabRequest, Response gitlab.GitlabResponse](client *GitlabClient, method string, host string, api string, request Request) Response {
	url := strings.TrimRight(host, "/") + api

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

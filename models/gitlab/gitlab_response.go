package gitlab

type GitlabResponse interface {
	OauthResponse | CreateUserResponse | CreatePersonalAccessTokenResponse |
		GetUserResponse | GetNamespaceResponse | CreateProjectResponse
}

package gitlab

type CreatePersonalAccessTokenRequest struct {
	Name      string   `json:"name"`
	ExpiresAt string   `json:"expires_at"`
	Scopes    []string `json:"scopes"`
}

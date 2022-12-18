package gitlab

import "time"

type CreatePersonalAccessTokenResponse struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Revoked    bool        `json:"revoked"`
	CreatedAt  time.Time   `json:"created_at"`
	Scopes     []string    `json:"scopes"`
	UserID     int         `json:"user_id"`
	LastUsedAt interface{} `json:"last_used_at"`
	Active     bool        `json:"active"`
	ExpiresAt  string      `json:"expires_at"`
	Token      string      `json:"token"`
}

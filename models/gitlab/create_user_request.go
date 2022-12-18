package gitlab

type CreateUserRequest struct {
	Admin    bool   `json:"admin"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

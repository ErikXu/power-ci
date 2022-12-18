package gitlab

import "time"

type CreateUserResponse struct {
	Id               int           `json:"id"`
	Username         string        `json:"username"`
	Name             string        `json:"name"`
	State            string        `json:"state"`
	AvatarURL        string        `json:"avatar_url"`
	WebURL           string        `json:"web_url"`
	CreatedAt        time.Time     `json:"created_at"`
	Bio              string        `json:"bio"`
	Location         interface{}   `json:"location"`
	PublicEmail      interface{}   `json:"public_email"`
	Skype            string        `json:"skype"`
	Linkedin         string        `json:"linkedin"`
	Twitter          string        `json:"twitter"`
	WebsiteURL       string        `json:"website_url"`
	Organization     interface{}   `json:"organization"`
	JobTitle         string        `json:"job_title"`
	Pronouns         interface{}   `json:"pronouns"`
	Bot              bool          `json:"bot"`
	WorkInformation  interface{}   `json:"work_information"`
	Followers        int           `json:"followers"`
	Following        int           `json:"following"`
	IsFollowed       bool          `json:"is_followed"`
	LocalTime        interface{}   `json:"local_time"`
	LastSignInAt     interface{}   `json:"last_sign_in_at"`
	ConfirmedAt      interface{}   `json:"confirmed_at"`
	LastActivityOn   interface{}   `json:"last_activity_on"`
	Email            string        `json:"email"`
	ThemeID          int           `json:"theme_id"`
	ColorSchemeID    int           `json:"color_scheme_id"`
	ProjectsLimit    int           `json:"projects_limit"`
	CurrentSignInAt  interface{}   `json:"current_sign_in_at"`
	Identities       []interface{} `json:"identities"`
	CanCreateGroup   bool          `json:"can_create_group"`
	CanCreateProject bool          `json:"can_create_project"`
	TwoFactorEnabled bool          `json:"two_factor_enabled"`
	External         bool          `json:"external"`
	PrivateProfile   bool          `json:"private_profile"`
	CommitEmail      string        `json:"commit_email"`
	IsAdmin          bool          `json:"is_admin"`
	Note             interface{}   `json:"note"`
	NamespaceID      int           `json:"namespace_id"`
	CreatedBy        struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"created_by"`
}

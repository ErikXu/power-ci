package gitlab

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	gitlabInitCmd.Flags().StringVarP(&Host, "host", "H", "http://example.com", "Gitlab Host")
	gitlabInitCmd.MarkFlagRequired("host")

	gitlabInitCmd.Flags().StringVarP(&User, "user", "u", "root", "Gitlab User")
	gitlabInitCmd.MarkFlagRequired("user")

	gitlabInitCmd.Flags().StringVarP(&Password, "password", "p", "", "Gitlab User Password")
	gitlabInitCmd.MarkFlagRequired("password")
}

var Host string
var User string
var Password string

var gitlabInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init gitlab",
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		gitlabClient := &GitlabClient{
			Client: *client,
			Host:   strings.TrimRight(Host, "/"),
		}

		response := gitlabClient.GrantOauthToken(User, Password)
		gitlabClient.AccessToken = response.AccessToken

		devopsUserId := 0
		users := gitlabClient.GetUserByUsername("devops_user")
		if len(users) >= 1 {
			devopsUserId = users[0].Id
		} else {
			devopsUser := gitlabClient.CreateUser(true, "devops_user", "Devops_User", "devops@example.com", "12345678")
			devopsUserId = devopsUser.Id
		}

		fmt.Println(devopsUserId)
	},
}

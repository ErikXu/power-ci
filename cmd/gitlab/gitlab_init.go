package gitlab

import (
	"fmt"
	"net/http"
	"power-ci/models/gitlab"

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
		}

		request := &gitlab.OauthRequest{}

		response := gitlabClient.GrantOauthToken("POST", Host, *request)

		fmt.Println(response.AccessToken)
	},
}

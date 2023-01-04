package gitlab

import (
	"fmt"
	"net/http"
	"power-ci/consts"
	"power-ci/utils"
	"strconv"
	"strings"

	"github.com/google/uuid"
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
		configs := utils.GetGitlabConfigs()

		client := &http.Client{}
		gitlabClient := &GitlabClient{
			Client: *client,
			Host:   strings.TrimRight(Host, "/"),
		}

		response := gitlabClient.GrantOauthToken("password", User, Password)
		gitlabClient.AccessToken = response.AccessToken

		devopsUserId := 0
		users := gitlabClient.GetUserByUsername("devops_user")
		if len(users) >= 1 {
			devopsUserId = users[0].Id
		} else {
			password := uuid.New()
			devopsUser := gitlabClient.CreateUser(true, consts.GitLabDefaultUser, "Devops_User", "devops@example.com", password.String(), true)
			devopsUserId = devopsUser.Id
			configs[consts.GitLabPasswordKey] = password.String()
		}

		var privateToken = gitlabClient.CreatePersonalAccessToken(devopsUserId, "devops_token", []string{"api"}, "2099-12-31")

		configs[consts.GitLabHostKey] = Host
		configs[consts.GitLabUserIdKey] = strconv.Itoa(devopsUserId)
		configs[consts.GitLabUserNameKey] = consts.GitLabDefaultUser
		configs[consts.GitLabPrivateTokenKey] = privateToken.Token

		path := utils.SaveGitlabConfigs(configs)

		fmt.Printf("Init success. Information was saved to [%s]\n", path)
	},
}

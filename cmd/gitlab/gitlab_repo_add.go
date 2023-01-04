package gitlab

import (
	"fmt"
	"net/http"
	"power-ci/consts"
	"power-ci/utils"
	"strings"

	"github.com/spf13/cobra"
)

var gitlabRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Gitlab repo tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gitlab repo tools")
	},
}

func init() {
	gitlabRepoCmd.AddCommand(gitlabRepoAddCmd)

	gitlabRepoAddCmd.Flags().StringVarP(&RepoGroup, "group", "g", "", "Repo group name")

	gitlabRepoAddCmd.Flags().StringVarP(&RepoName, "name", "n", "", "Repo name")
	gitlabRepoAddCmd.MarkFlagRequired("name")

	gitlabRepoAddCmd.Flags().StringVarP(&RepoType, "type", "t", "", "Repo type, eg: go, dotnet, java...")
	gitlabRepoAddCmd.MarkFlagRequired("type")

	gitlabRepoAddCmd.Flags().StringVarP(&RepoOwner, "owner", "o", "", "Repo owner")
}

var RepoGroup string
var RepoName string
var RepoType string
var RepoOwner string

var gitlabRepoAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add gitlab repo",
	Run: func(cmd *cobra.Command, args []string) {
		configs := utils.GetGitlabConfigs()

		client := &http.Client{}
		gitlabClient := &GitlabClient{
			Client:       *client,
			Host:         strings.TrimRight(configs[consts.GitLabHostKey], "/"),
			PrivateToken: configs[consts.GitLabPrivateTokenKey],
		}

		namespaces := gitlabClient.GetNamespaces()

		namespaceId := 0
		if RepoGroup == "" {
			for _, namespace := range namespaces {
				if namespace.Kind == "user" && namespace.FullPath == consts.GitLabDefaultUser {
					namespaceId = namespace.Id
				}
			}
		}

		project := gitlabClient.CreateProject(RepoName, namespaceId)

		commit := gitlabClient.CreateCommit(project.Id, "main", "create", "README.md", "# Try add readme", "Try Add readme.")
		fmt.Println(commit)
	},
}

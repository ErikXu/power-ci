package gitlab

import (
	"fmt"

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

	gitlabRepoAddCmd.Flags().StringVarP(&RepoNamespace, "group", "g", "", "Repo group name")

	gitlabRepoAddCmd.Flags().StringVarP(&RepoName, "name", "n", "", "Repo name")
	gitlabRepoAddCmd.MarkFlagRequired("name")

	gitlabRepoAddCmd.Flags().StringVarP(&RepoType, "type", "t", "", "Repo type, eg: go, dotnet, java...")
	gitlabRepoAddCmd.MarkFlagRequired("type")

	gitlabRepoAddCmd.Flags().StringVarP(&RepoOwner, "owner", "o", "", "Repo owner")
}

var RepoNamespace string
var RepoName string
var RepoType string
var RepoOwner string

var gitlabRepoAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add gitlab repo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add gitlab repo")
	},
}

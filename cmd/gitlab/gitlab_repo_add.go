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

var goIgnore = `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/`

var goDockerfile = `# 使用 alpine 作为基础镜像
FROM alpine:3.14

# 设置 Gin Mode 为 release
ENV GIN_MODE release

# 拷贝可执行文件到 /app
COPY . /app

# 设置 /app 为工作目录
WORKDIR /app

# 操作系统使用阿里云的源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装 tini
RUN apk update
RUN apk add tini

# 暴漏 8080 端口
EXPOSE 8080

# 使用 tini 模式启动程序
ENTRYPOINT ["tini", "--"]
CMD ["/app/app"]`

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

		gitlabClient.CreateCommit(project.Id, "main", "create", ".gitignore", goIgnore, "Add .gitignore.")
		gitlabClient.CreateCommit(project.Id, "main", "create", "docker/Dockerfile", goDockerfile, "Add Dockerfile.")
	},
}

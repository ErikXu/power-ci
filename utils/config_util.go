package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"power-ci/consts"
)

var configName = "power-ci.json"
var gitlabConfigName = "gitlab.json"

func GetGitlabConfigs() map[string]string {
	return getConfigs(gitlabConfigName)
}

func SaveGitlabConfigs(configs map[string]string) string {
	return setConfigs(gitlabConfigName, configs)
}

func GetConfigs() map[string]string {
	return getConfigs(configName)
}

func SaveConfigs(configs map[string]string) string {
	return setConfigs(configName, configs)
}

func getConfigs(filename string) map[string]string {
	bytes := readConfig(filename)
	if bytes == nil {
		return make(map[string]string)
	}

	result := make(map[string]string)
	json.Unmarshal(bytes, &result)
	return result
}

func setConfigs(filename string, configs map[string]string) string {
	bytes, _ := json.MarshalIndent(configs, "", "    ")
	return saveConfig(filename, bytes)
}

func readConfig(filename string) []byte {
	homeDir, _ := os.UserHomeDir()
	os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

	filepath := path.Join(homeDir, consts.Workspace, filename)
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Cannot read file [%s]", filepath)
		}
		return nil
	}
	return bytes
}

func saveConfig(filename string, bytes []byte) string {
	homeDir, _ := os.UserHomeDir()
	os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

	filepath := path.Join(homeDir, consts.Workspace, filename)
	f, _ := os.Create(filepath)
	f.Write(bytes)

	return filepath
}

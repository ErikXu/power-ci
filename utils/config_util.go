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

func GetConfigs() map[string]string {
	return getConfigs(configName)
}

func SaveConfigs(configs map[string]string) {
	setConfigs(configName, configs)
}

func getConfigs(filename string) map[string]string {
	bytes := readConfig(filename)
	if bytes == nil {
		return nil
	}

	result := make(map[string]string)
	json.Unmarshal(bytes, &result)
	return result
}

func setConfigs(filename string, configs map[string]string) {
	bytes, _ := json.MarshalIndent(configs, "", "    ")
	saveConfig(filename, bytes)
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

func saveConfig(filename string, bytes []byte) {
	homeDir, _ := os.UserHomeDir()
	os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

	filepath := path.Join(homeDir, consts.Workspace, filename)
	f, _ := os.Create(filepath)
	f.Write(bytes)
}

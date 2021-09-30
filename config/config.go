package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	homeDir = "/usr/local/gogen"
	// 即将废弃
	PluginsDir2 = homeDir + "/plugins"
)

func createDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, 0777)
	}
	return nil
}

func createFile() error {
	return nil
}

func InitGoGen() error {
	dirList := []string{UserHomeDir(), PluginsDir()}
	for _, dir := range dirList {
		if err := createDir(dir); err != nil {
			return err
		}
	}
	return nil
}

func UserHomeDir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		log.Println("UserHomeDir error:", err)
		return homeDir
	}
	return h + "/gogen"
}

func PluginsDir() string {
	return fmt.Sprintf("%s/%s", UserHomeDir(), "plugins")
}

type Plugin struct {
	AppName    string
	PluginName string
}
type pluginJson []*Plugin

func getPluginsJson() pluginJson {
	c := path.Join(PluginsDir(), "plugins.json")
	b, _ := ioutil.ReadFile(c)
	var list pluginJson
	if len(b) > 0 {
		_ = json.Unmarshal(b, &list)
	}
	return list
}

func SavePluginsJson(data *Plugin) error {
	c := path.Join(PluginsDir(), "plugins.json")
	list := getPluginsJson()
	b2, _ := json.Marshal(append(list, data))
	return ioutil.WriteFile(c, b2, os.ModePerm)
}

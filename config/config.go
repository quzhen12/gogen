package config

import "os"

const (
	homeDir    = "/usr/local/gogen"
	PluginsDir = homeDir + "/plugins"

	pluginsJson = PluginsDir + "/plugins.json"
)

func createDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, os.ModePerm)
	}
	return nil
}

func createFile() error {
	return nil
}

func InitGoGen() error {
	dirList := []string{homeDir, PluginsDir}
	for _, dir := range dirList {
		if err := createDir(dir); err != nil {
			return err
		}
	}
	return nil
}

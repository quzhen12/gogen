package config

import (
	"fmt"
	"log"
	"os"
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

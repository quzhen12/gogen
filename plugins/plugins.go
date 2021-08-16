package plugins

import (
	"encoding/json"
	"fmt"
	"gogen/config"
	"gogen/pkg/common"
	"io/ioutil"
	"path"

	"github.com/spf13/viper"
)

type PluginOpt interface {
	Install(src string) error
	Remove(src string) error
	Clean() error
}

type pluginsInfo struct {
	AppName     string `mapstructure:"app_name"`
	ProjectName string `mapstructure:"project_name"`
}

type plugins struct {
	info *pluginsInfo
}

func NewPlugins() PluginOpt {
	return &plugins{}
}

func (p *plugins) Install(src string) error {
	err := p.loadPluginsConfig(src)
	if err != nil {
		return err
	}
	err = common.CopyFiles(config.PluginsDir, src)
	if err != nil {
		return err
	}
	return savePluginsConfig(map[string]interface{}{
		"app_name": p.info.AppName,
	})
}

func (p *plugins) loadPluginsConfig(src string) error {
	viper.AddConfigPath(src)
	viper.SetConfigName("gogen") // name of config file (without extension)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return err
	}
	info := &pluginsInfo{}
	err = viper.Unmarshal(info)
	if err != nil {
		return err
	}
	p.info = info
	fmt.Println("name", viper.GetString("app_name"))
	return nil
}

func (p *plugins) Remove(src string) error {
	panic("implement me")
}

func (p *plugins) Clean() error {
	panic("implement me")
}

func savePluginsConfig(data map[string]interface{}) error {
	c := path.Join(config.PluginsDir, "plugins.json")
	b, _ := ioutil.ReadFile(c)

	cfg := []map[string]interface{}{}
	if len(b) > 0 {
		_ = json.Unmarshal(b, &cfg)
	}
	cfg = append(cfg, data)
	bb, _ := json.Marshal(cfg)
	ioutil.WriteFile(c, bb, 0700)
	return nil
}

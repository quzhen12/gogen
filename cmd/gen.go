package cmd

import (
	"fmt"
	"gogen/config"
	"gogen/gen"
	"gogen/plugins"
	"path"

	"github.com/quzhen12/plugins/file"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var appName = ""
var pluginsName = ""

func init() {
	rootCmd.PersistentFlags().StringVar(&appName, "app_name", "", "app name")
	rootCmd.PersistentFlags().StringVar(&pluginsName, "plugins_name", "", "app name")
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new project",
	Long:  "Generate a project by template that's installed.",
	Run: func(cmd *cobra.Command, args []string) {
		if appName == "" || pluginsName == "" {
			return
		}
		genProject()
		fmt.Println("Success!")
	},
}

func genProject() {
	p := plugins.NewPlugins()
	pcfs, err := p.GetPluginsConfig()
	if err != nil {
		return
	}
	g := gen.NewGentor()
	gg := gen.NewGen(g)
	gg.G.SetProjectName(appName)
	gg.G.SetOldProjectName(pcfs[pluginsName].ProjectName)
	err = file.Mkdir(appName)
	if err != nil {
		zap.L().Error("mkdir appName", zap.Any("err", err))
		return
	}
	err = gg.Travels(path.Join(config.PluginsDir(), pluginsName))
	if err != nil {
		zap.L().Error("travel", zap.Any("err", err))
	}
}

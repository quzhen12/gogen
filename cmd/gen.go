package cmd

import (
	"fmt"
	"gogen/config"
	"gogen/gen"
	"gogen/pkg/common"
	"gogen/plugins"
	"path"

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
		fmt.Println("Success!")
		genProject()
	},
}

func genProject() {
	p := plugins.NewPlugins()
	pcfs, err := p.GetPluginsConfig()
	if err != nil {
		return
	}
	if appName == "" {
		return
	}
	g := gen.NewGentor()
	gg := gen.NewGen(g)
	gg.G.SetProjectName(appName)
	gg.G.SetOldProjectName(pcfs[pluginsName].ProjectName)
	err = common.Mkdir(appName)
	if err != nil {
		fmt.Println("mkdir err", err)
		return
	}
	err = gg.Travels(path.Join(config.PluginsDir, pluginsName))
	if err != nil {
		fmt.Println(err)
	}
}

package cmd

import (
	"fmt"
	"gogen/config"
	"gogen/gen"
	"gogen/plugins"
	"os"
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
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		gogen()
	},
}

func gogen() {
	err := config.InitGoGen()
	if err != nil {
		fmt.Println("err", err)
	}
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
	err = os.Mkdir("./"+appName, 0777)
	if err != nil {
		fmt.Println("mkdir err", err)
	}
	err = gg.Travels(path.Join(config.PluginsDir, pluginsName))
	if err != nil {
		fmt.Println(err)
	}
}

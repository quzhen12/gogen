package cmd

import (
	"flag"
	"fmt"
	"gogen/config"
	"gogen/gen"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func init() {
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
	flag.String("install", "", "install")
	name := flag.String("name", "", "create")
	pkg := flag.String("pkg", "", "pkg")
	p := flag.String("p", "", "plugins")
	flag.Parse()
	fmt.Println("name: ", *name, *p)
	if *name == "" || *pkg == "" || *p == "" {
		return
	}
	g := gen.NewGentor()
	gg := gen.NewGen(g)
	gg.G.SetProjectName(*name)
	gg.G.SetOldProjectName(*pkg)
	err = os.Mkdir("./"+*name, 0777)
	if err != nil {
		fmt.Println("mkdir err", err)
	}
	err = gg.Travels(path.Join(config.PluginsDir, *p))
	if err != nil {
		fmt.Println(err)
	}
}

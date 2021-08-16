package cmd

import (
	"fmt"
	"gogen/plugins"

	"github.com/spf13/cobra"
)

var pkgSrc string

func init() {
	rootCmd.PersistentFlags().StringVar(&pkgSrc, "pkg", "", "pkg src")
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		p := plugins.NewPlugins()
		fmt.Println("pkg src", pkgSrc)
		if pkgSrc == "" {
			return
		}
		err := p.Install(pkgSrc)
		if err != nil {
			fmt.Println("install error,", err)
		}
	},
}

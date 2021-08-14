package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"gogen/config"

	"gogen/gen"
)

var (
	appDir     = "/usr/local/gogen"
	pluginsDir = path.Join(appDir, "plugins")
)

func main() {
	err := config.InitGoGen()
	if err != nil {
		fmt.Println("err", err)
	}
	flag.String("install", "", "install")
	name := flag.String("name", "", "create")
	pkg := flag.String("pkg", "", "pkg")
	flag.Parse()
	fmt.Println("name: ", *name)
	if *name == "" || *pkg == "" {
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
	err = gg.Travels("temp/tmpweb")
	if err != nil {
		fmt.Println(err)
	}
}

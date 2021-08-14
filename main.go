package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.String("install", "", "install")
	name := flag.String("name", "", "create")
	pkg := flag.String("pkg", "", "pkg")
	flag.Parse()
	fmt.Println("name: ", *name)
	if *name == "" || *pkg == "" {
		return
	}
	g := NewGentor()
	gg := NewGen(g)
	gg.g.SetProjectName(*name)
	gg.g.SetOldProjectName(*pkg)
	err := os.Mkdir("./"+*name, 0777)
	if err != nil {
		fmt.Println("mkdir err", err)
	}
	err = gg.travels("temp/tmpweb")
	if err != nil {
		fmt.Println(err)
	}
}

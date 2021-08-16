package main

import (
	"fmt"
	"gogen/cmd"
	"gogen/config"
)

func main() {
	err := config.InitGoGen()
	if err != nil {
		fmt.Println("err", err)
	}
	cmd.Execute()
}

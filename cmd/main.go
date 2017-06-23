package main

import (
	"fmt"
	"os"

	"github.com/tiabc/jobrunner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n    jobrunner <config-file>")
		return
	}
	conf, err := jobrunner.NewConfigFromFile(os.Args[1])
	fmt.Println(conf, err)
}

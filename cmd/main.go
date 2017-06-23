package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tiabc/jobrunner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n    jobrunner <config-file>")
		return
	}
	r, err := jobrunner.NewFromFile(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to parse config %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}
	ctx, _ := context.WithCancel(context.Background())
	r.Run(ctx)
	// TODO: Graceful shutdown.
}

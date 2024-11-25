package main

import (
	"fmt"
	"os"

	"github.com/la-plas-growth/go-docker-ai/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

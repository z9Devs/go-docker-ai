package main

import (
	"fmt"
	"os"

	"github.com/la-plas-growth/GO-DockerLint-AI/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

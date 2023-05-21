package main

import (
	"fmt"
	"go-distributed-system/core"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing argument")
		fmt.Println("Usage: go run main.go <argument>")
		os.Exit(1) // Exit with non-zero status to indicate an error
	}
	nodeType := os.Args[1]
	switch nodeType {
	case "master":
		core.GetMasterNode().Start()
	case "worker":
		core.GetWorkerNode().Start()
	default:
		panic("Error: invalid node type in argument!")
	}
}

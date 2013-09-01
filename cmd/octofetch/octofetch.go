package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("USAGE: octofetch [resourceId]")
}

// fetch whatever
func main() {
	args := os.Args
	if len(args) != 1 {
		usage()
		os.Exit(1)
	}

}

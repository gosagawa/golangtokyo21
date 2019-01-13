package main

import (
	"fmt"
	"os"

	"github.com/gosagawa/golangtokyo21/tree"
)

func main() {

	args := os.Args

	if len(args) != 2 {
		usage()
		os.Exit(2)
	}

	dir := args[1]
	err := tree.OutputTree(dir, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func usage() {

	usage := ""
	usage += "Usage: tree: dir_pass\n"
	_, err := fmt.Fprintf(os.Stderr, usage)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

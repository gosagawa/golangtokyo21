package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gosagawa/golangtokyo21/tree"
)

func main() {

	var (
		deps = flag.Int("L", 0, "階層の深さ")
	)

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		usage()
		os.Exit(2)
	}

	dir := args[0]
	tree, err := tree.NewTree(*deps)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	err = tree.OutputTree(dir, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func usage() {

	usage := ""
	usage += "Usage: tree: tree [OPTION] dir_pass\n"
	usage += "  -L deps\n"
	usage += "      階層の深さ。0(デフォルト)で無制限 \n"
	_, err := fmt.Fprintf(os.Stderr, usage)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

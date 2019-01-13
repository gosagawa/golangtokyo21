package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gosagawa/golangtokyo21/tree"
)

func main() {

	var (
		deps       = flag.Int("L", 0, "階層の深さ")
		dirOnly    = flag.Bool("D", false, "ディレクトリのみ表示")
		maskGopher = flag.Bool("G", false, "ファイル名をGopherでマスク")
	)

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		usage()
		os.Exit(2)
	}

	dir := args[0]
	option := tree.Option{
		Deps:       *deps,
		DirOnly:    *dirOnly,
		MaskGopher: *maskGopher,
	}
	tree, err := tree.NewTree(option)
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

	usage := `
Usage: tree: tree [OPTION] dir_pass
    -L deps
        階層の深さ。0(デフォルト)で無制限
    -D 
        ディレクトリのみ表示
    -G 
        ファイル名をGopherでマスク
`
	_, err := fmt.Fprintf(os.Stderr, usage)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

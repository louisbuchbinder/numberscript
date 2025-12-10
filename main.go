package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/louisbuchbinder/numberscript/lib"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		panic(fmt.Errorf("usage numberscript <filename>"))
	}
	filename := flag.Args()[0]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read file:", err)
		os.Exit(1)
	}

	b, err := lib.NewInterpreter().Exec([]rune(string(content)))
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}

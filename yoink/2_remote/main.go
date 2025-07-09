package main

import (
	"context"
	"fmt"
	"github.com/tedla-brandsema/yoink"
	"os"
)

func main() {
	// Open the root file
	name := "./data/sonnet-18-remote.txt"
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Resolve .yoink commands in the root file
	txt, err := yoink.Parse(context.Background(), file, name)
	if err != nil {
		panic(err)
	}
	fmt.Println(txt)
}

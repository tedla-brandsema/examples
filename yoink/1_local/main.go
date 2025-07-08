package main

import (
	"context"
	"fmt"
	"github.com/tedla-brandsema/yoink"
	"os"
)

func main() {

	path := "./data/sonnet-18.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	txt, err := yoink.Parse(context.Background(), file, path)
	if err != nil {
		panic(err)
	}
	fmt.Println(txt)
}

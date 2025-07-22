package main

import (
	"context"
	"fmt"
	"github.com/tedla-brandsema/yoink"
	"os"
	"strings"
	"sync"
)

// START Parser OMIT
type CountParser struct {
	mut   sync.Mutex
	count int
}

func (p *CountParser) Parse(fileName string, lineNumber int, inputLine string) (string, error) {
	// Since we share state over multiple goroutines, we need to guard against possible race conditions
	p.mut.Lock()
	defer p.mut.Unlock()

	// Increment the counter
	p.count++

	// Return the invocation count
	return fmt.Sprintf("Command %q has been invoked %d times", strings.Fields(inputLine)[0], p.count), nil
}

// END Parser OMIT

func main() {
	// Register an instance of CountParser
	// START RegisterParser OMIT
	yoink.RegisterParser("count", &CountParser{})
	// END RegisterParser OMIT

	// Open the root file
	name := "./data/count.txt"
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Resolve .count commands in the root file
	txt, err := yoink.Parse(context.Background(), file, name)
	if err != nil {
		panic(err)
	}
	fmt.Println(txt)
}

package main

import (
	"context"
	"fmt"
	"github.com/tedla-brandsema/yoink"
	"os"
	"strings"
)

// START ParserFunc OMIT
func HelloParser(sourceFile string, sourceLine int, cmd string) (string, error) {
	// Default subject of the greeting
	subject := fmt.Sprintf("from %s", sourceFile)

	// Split the command into its parts where the first part is always the command associated with this parser.
	// In this case .hello
	parts := strings.Fields(cmd)
	if len(parts) > 1 {
		// Substitute the default subject with the arguments provided
		subject = strings.Join(parts[1:], " ")
	}

	// Return the greeting. Starting with the line number, followed by the greeting to our subject.
	return fmt.Sprintf("%d. Hello, %s!", sourceLine, subject), nil
}

// END ParserFunc OMIT

func main() {
	// Register the HelloParser, which is a ParserFunc
	// START RegisterParserFunc OMIT
	yoink.RegisterParserFunc("hello", HelloParser)
	// END RegisterParserFunc OMIT

	// Open the root file
	name := "./data/hello.txt"
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

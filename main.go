package main

import (
	_ "embed"
	"log"

	"github.com/skelouse/code-to-gpt/parser"
)

func main() {
	if err := parser.Run(); err != nil {
		log.Fatal(err)
	}
}

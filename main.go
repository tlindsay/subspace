package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tlindsay/subspace/subspace"
	"github.com/tlindsay/subspace/api"
	"github.com/tlindsay/subspace/tui"
)

func main() {
	var numParagraphs int
	var numLines int
	var character string
	var interactive bool
	var shouldPrintCharacters bool
	var shouldStartServer bool
	var port int

	flag.IntVar(&numParagraphs, "paragraphs", 1, "the number of paragraphs to print")
	flag.IntVar(&numParagraphs, "p", 1, "the number of paragraphs to print")
	flag.IntVar(&numLines, "lines", 3, "the number of lines to print")
	flag.IntVar(&numLines, "l", 3, "the number of lines to print")
	flag.StringVar(&character, "c", "picard", "the character whose dialog you want")
	flag.StringVar(&character, "character", "picard", "the character whose dialog you want")
	flag.BoolVar(&shouldPrintCharacters, "lc", false, "list the available characters")
	flag.BoolVar(&shouldPrintCharacters, "list-chars", false, "list the available characters")
	flag.BoolVar(&shouldStartServer, "s", false, "start JSON server")
	flag.BoolVar(&shouldStartServer, "serve", false, "start JSON server")
	flag.BoolVar(&interactive, "i", false, "start interactive TUI")
	flag.BoolVar(&interactive, "interactive", false, "start interactive TUI")
	flag.IntVar(&port, "port", 1701, "the port to serve on")
	flag.Parse()

	if (interactive) {
		tui.StartTUI()
		os.Exit(0)
	}

	if shouldPrintCharacters {
		chars := subspace.ListAllCharacters()
		fmt.Println(strings.Join(chars, "\n"))
		return
	}

	if shouldStartServer {
		api.StartServer(port)
	}

	text, err := subspace.MakeItSo(numParagraphs, numLines, character)
	if err != nil {
		fmt.Println("Fatal Error:", err)
		os.Exit(1)
	}
	fmt.Println(strings.Join(text, "\n"))
}

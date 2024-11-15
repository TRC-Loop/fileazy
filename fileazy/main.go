package main

import (
	"fileazy/assets"
	"flag"
	"fmt"
)

func main() {
	// Flags
	file := flag.String("f", "", assets.FileHelp)
	noOpen := flag.Bool("no", false, assets.NoOpenHelp)
	openConfig := flag.Bool("c", false, assets.OpenConfigHelp)
	searchForFile := flag.String("s", "", assets.SearchForFileHelp)

	// Parse Flags
	flag.Parse()

	// Print Flags (for Debugging)
	fmt.Printf("Flags: file: %s, noOpen: %t, openConfig: %t searchForFile: %s\n", *file, *noOpen, *openConfig, *searchForFile)

	// TODO: Implement Configmenu, Fuzzysearch(Menu), Openmenu.
}

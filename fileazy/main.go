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
	editConfig := flag.Bool("ec", false, assets.EditConfigDirectlyHelp)

	// Parse Flags
	flag.Parse()

	// Print Flags (for Debugging)
	fmt.Printf("Flags: file: %s, noOpen: %t, openConfig: %t searchForFile: %s ec: %t\n", *file, *noOpen, *openConfig, *searchForFile, *editConfig)

	// TODO: Implement Configmenu, Fuzzysearch(Menu), Openmenu.
}

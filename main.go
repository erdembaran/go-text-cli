package main

import (
	"flag"
	"fmt"
)

func countCharacters(text string) int {
	return len(text)
}

func countCharactersWithoutSpaces(text string) int {
	count := 0

	for _, char := range text {
		if char != ' ' {
			count++
		}
	}
	return count
}

func main() {

	// Define flags
	allCharsCmd := flag.Bool("all", false, "Count all characters including spaces")
	noSpaceCmd := flag.Bool("nospace", false, "Count all characters excluding spaces")

	// Parse the flags
	flag.Parse()

	// Get the input text from remaining arguments
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide text to analyze")
        fmt.Println("Usage:")
        fmt.Println("  go run main.go -all \"your text\"    (count all characters)")
        fmt.Println("  go run main.go -nospace \"your text\" (count excluding spaces)")
        return
	}

	inputText := flag.Args()[0]

	// Handle different commands
	if *allCharsCmd {
		count := countCharacters(inputText)
		fmt.Printf("Total characters (including spaces): %d\n", count)
	} else if *noSpaceCmd {
		count := countCharactersWithoutSpaces(inputText)
		fmt.Printf("Total characters (excluding spaces): %d\n", count)
	} else {
        fmt.Println("Please specify a command: -all or -nospace")
        fmt.Println("Example: go run main.go -all \"Hello World\"")
    }
}

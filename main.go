package main

import (
	"flag"
	"fmt"
	"strings"
)

// showHelp prints detailed usage information
func showHelp() {
	fmt.Println(`
	Text Processor - A tool for analyzing text
	
	Commands:
		-all        Count all characters including spaces
		-nospace    Count characters excluding spaces
		-words      Count total words
		-help       Show this help message
	
	Usage:
		go run main.go [command] "your text"
	
	Examples:
		go run main.go -all "Hello World"
			Count all characters including spaces
		
		go run main.go -nospace "Hello World"
			Count characters excluding spaces

		go run main.go -words "Hello World"
            Count total words
	
	Note:
		Text should be enclosed in quotes if it contains spaces
		Only one command can be used at a time
	`)
}

func countCharacters(text string) int {
	return len(text)
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
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
	noSpaceCmd := flag.Bool("nospace", false, "Count characters excluding spaces")
	wordsCmd := flag.Bool("words", false, "Count total words")
	helpCmd := flag.Bool("help", false, "Show help message")

	// Parse the flags
	flag.Parse()

	// Check for help command first
	if *helpCmd {
		showHelp()
		return
	}

	// Check for input text
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide text to analyze")
		fmt.Println("Use -help for more information")
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
	} else if *wordsCmd {
        count := countWords(inputText)
        fmt.Printf("Total words: %d\n", count)
    } else {
		fmt.Println("Please specify a command. Use -help to see available commands")
	}
}

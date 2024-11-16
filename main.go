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
		-wordlen    Show longest and shortest words
		-readtime   Estimate reading time (based on 225 WPM)
		-help       Show this help message
	
	Usage:
		go run main.go [command] "your text"
	
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

func findWordExtremes(text string) (longest, shortest string) {
	words := strings.Fields(text)

	if len(words) == 0 {
        return "", ""
    }

	longest = words[0]
	shortest = words[0]

	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
		if len(word) < len(shortest) {
			shortest = word
		}
	}

	return longest, shortest
}

func readingEstimationTime(wordCount int) float64 {

	var averageReadWPM float64 = 225

	var readingTimeInMinutes float64 = float64(wordCount) / float64(averageReadWPM)

	return readingTimeInMinutes
}

func main() {

	// Define flags
	allCharsCmd := flag.Bool("all", false, "Count all characters including spaces")
	noSpaceCmd := flag.Bool("nospace", false, "Count characters excluding spaces")
	wordsCmd := flag.Bool("words", false, "Count total words")
	wordExtremeCmd := flag.Bool("wordlen", false, "Find longest and shortest words")
	readingEstimation := flag.Bool("readtime", false, "Estimate the reading time")
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
    } else if *wordExtremeCmd {
		longest, shortest := findWordExtremes(inputText)
		if longest == "" && shortest == "" {
			fmt.Println("No words found in input")
		} else {
			fmt.Printf("Longest word (%d chars): %s\n", len(longest), longest)
			fmt.Printf("Shortest word (%d chars): %s\n", len(shortest), shortest)
		}
	} else if *readingEstimation {
		count := countWords(inputText)
		readingTime := readingEstimationTime(count)
		
		if readingTime < 1 {
			fmt.Printf("Average reading time: %.0f seconds\n", readingTime*60)
		} else if readingTime == 1 {
			fmt.Println("Average reading time: 1 minute")
		} else {
			minutes := int(readingTime)
			seconds := int((readingTime - float64(minutes)) * 60)
			
			if seconds >= 1 {
				fmt.Printf("Average reading time: %d minutes %d seconds\n", minutes, seconds)
			} else {
				fmt.Printf("Average reading time: %d minutes\n", minutes)
			}

		}
	} else {
		fmt.Println("Please specify a command. Use -help to see available commands")
	}
}

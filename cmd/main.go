package main

import (
	"flag"
	"fmt"
	"strings"
)

func generateHeader(title string, lang string) string {
	// Maximum line width for the header
	const lineWidth = 60

	// Clean and convert the title to uppercase
	title = strings.TrimSpace(title)
	title = strings.ToUpper(title) // Convert title to uppercase

	// Calculate padding to center the title
	totalPadding := lineWidth - len(title) - 2 // subtract 2 for the characters on both ends
	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding
	leftPadStr := strings.Repeat(" ", leftPadding)
	rightPadStr := strings.Repeat(" ", rightPadding)

	if totalPadding < 0 {
		title = title[:lineWidth-2]
		leftPadStr = ""
		rightPadStr = ""
	}

	// Build the header based on the language
	var header strings.Builder
	switch lang {
	case "go", "java", "rust", "ts":
		border := strings.Repeat("*", lineWidth)
		header.WriteString(fmt.Sprintf("//%s\n", border))
		header.WriteString(fmt.Sprintf("//*%s%s%s*\n", leftPadStr, title, rightPadStr))
		header.WriteString(fmt.Sprintf("//%s", border))
	case "python":
		border := strings.Repeat("#", lineWidth)
		header.WriteString(fmt.Sprintf("#%s\n", border))
		header.WriteString(fmt.Sprintf("#%s%s%s#\n", leftPadStr, title, rightPadStr))
		header.WriteString(fmt.Sprintf("#%s", border))
	}

	return header.String()
}

func main() {
	// Define flags for different languages
	goFlag := flag.Bool("go", false, "Generate header for Go")
	javaFlag := flag.Bool("java", false, "Generate header for Java")
	pythonFlag := flag.Bool("python", false, "Generate header for Python")
	rustFlag := flag.Bool("rust", false, "Generate header for Rust")
	typescriptFlag := flag.Bool("ts", false, "Generate header for TypeScript")
	flag.Parse()

	// Get the title from command arguments
	if flag.NArg() < 1 {
		fmt.Println("Please provide a title for the header")
		fmt.Println("Example: hd-gen -go \"Variable Declaration\"")
		return
	}

	title := strings.Join(flag.Args(), " ")

	// Check the selected flag and generate the corresponding header
	switch {
	case *goFlag:
		result := generateHeader(title, "go")
		fmt.Println(result)
	case *javaFlag:
		result := generateHeader(title, "java")
		fmt.Println(result)
	case *pythonFlag:
		result := generateHeader(title, "python")
		fmt.Println(result)
	case *rustFlag:
		result := generateHeader(title, "rust")
		fmt.Println(result)
	case *typescriptFlag:
		result := generateHeader(title, "ts")
		fmt.Println(result)
	default:
		fmt.Println("Please specify a language:")
		fmt.Println("  -go: for Go")
		fmt.Println("  -java: for Java")
		fmt.Println("  -python: for Python")
		fmt.Println("  -rust: for Rust")
		fmt.Println("  -ts: for TypeScript")
	}
}

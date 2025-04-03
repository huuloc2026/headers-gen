package main

import (
	"flag"
	"fmt"
	"strings"
)

func generateHeader(title string, lang string) string {
	// Tính độ dài tối đa của header
	const lineWidth = 60

	// Chuẩn hóa title và chuyển sang chữ in hoa
	title = strings.TrimSpace(title)
	title = strings.ToUpper(title) // Chuyển title sang chữ in hoa

	// Tính padding để title căn giữa
	totalPadding := lineWidth - len(title) - 2 // trừ đi 2 cho ký tự ở hai đầu
	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding
	leftPadStr := strings.Repeat(" ", leftPadding)
	rightPadStr := strings.Repeat(" ", rightPadding)

	if totalPadding < 0 {
		title = title[:lineWidth-2]
		leftPadStr = ""
		rightPadStr = ""
	}

	// Xây dựng header theo ngôn ngữ
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
	// Định nghĩa flags cho các ngôn ngữ
	goFlag := flag.Bool("go", false, "Generate header for Go")
	javaFlag := flag.Bool("java", false, "Generate header for Java")
	pythonFlag := flag.Bool("python", false, "Generate header for Python")
	rustFlag := flag.Bool("rust", false, "Generate header for Rust")
	typescriptFlag := flag.Bool("ts", false, "Generate header for TypeScript")
	flag.Parse()

	// Lấy title từ arguments
	if flag.NArg() < 1 {
		fmt.Println("Vui lòng cung cấp tiêu đề cho header")
		fmt.Println("Ví dụ: hdgen -go \"Variable Declaration\"")
		return
	}

	title := strings.Join(flag.Args(), " ")

	// Kiểm tra và tạo header dựa trên flag được chọn
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
		fmt.Println("Vui lòng chỉ định ngôn ngữ:")
		fmt.Println("  -go: for Go")
		fmt.Println("  -java: for Java")
		fmt.Println("  -python: for Python")
		fmt.Println("  -rust: for Rust")
		fmt.Println("  -ts: for TypeScript")
	}
}

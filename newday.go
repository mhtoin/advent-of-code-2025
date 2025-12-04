//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

const dayTemplate = `package main

import (
	"fmt"

	"github.com/mhtoin/advent-of-code-2025/common"
)

func main() {
	solvePart1()
	solvePart2()
}

func solvePart1() {
	total := 0

	common.ForEachLine({{.Day}}, func(line string) {
		// TODO: implement
		_ = line
	})

	fmt.Printf("Part 1: %d\n", total)
}

func solvePart2() {
	total := 0

	common.ForEachLine({{.Day}}, func(line string) {
		// TODO: implement
		_ = line
	})

	fmt.Printf("Part 2: %d\n", total)
}
`

type templateData struct {
	Day int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run newday.go <day>")
		fmt.Println("Example: go run newday.go 4")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > 25 {
		fmt.Println("Error: day must be a number between 1 and 25")
		os.Exit(1)
	}

	dirName := fmt.Sprintf("day%02d", day)
	dirPath := filepath.Join(".", dirName)

	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		fmt.Printf("Error: directory %s already exists\n", dirName)
		os.Exit(1)
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	tmpl := template.Must(template.New("day").Parse(dayTemplate))
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, templateData{Day: day}); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		os.Exit(1)
	}

	mainPath := filepath.Join(dirPath, "main.go")
	if err := os.WriteFile(mainPath, buf.Bytes(), 0644); err != nil {
		fmt.Printf("Error creating main.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Created %s/main.go\n", dirName)
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("  1. Run: go run ./%s\n", dirName)
	fmt.Printf("     (This will auto-download the input file)\n")
	fmt.Printf("  2. Implement solvePart1() and solvePart2()\n")
}

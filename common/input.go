package common

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// Load .env file (silently ignore if not found)
	_ = godotenv.Load()
}

func DownloadInputFile(day int) string {
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", day)

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		panic("AOC_SESSION environment variable not set")
	}

	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Failed to download input file: %s", resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	check(err)
	dayStr := fmt.Sprintf("day%02d", day)
	dirPath := filepath.Join("./", dayStr)
	err = os.MkdirAll(dirPath, os.ModePerm)
	check(err)

	filePath := filepath.Join(dirPath, "input.txt")
	err = os.WriteFile(filePath, body, 0644)
	check(err)
	fmt.Printf("Input file for day %d downloaded successfully.\n", day)

	return filePath

}

func GetInputFile(day int) (*os.File, error) {
	dayStr := fmt.Sprintf("day%02d", day)
	path := filepath.Join("./", dayStr, "input.txt")

	fmt.Printf("Looking for input file at: %s\n", path)

	file, err := os.Open(path)

	if os.IsNotExist(err) {
		fmt.Printf("Input file for day %d not found. Downloading...\n", day)
		filePath := DownloadInputFile(day)
		file, err = os.Open(filePath)
	}
	fmt.Printf("Input file for day %d opened successfully.\n", day)
	return file, err
}

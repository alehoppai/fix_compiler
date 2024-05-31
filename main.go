package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
    "github.com/fix_compiler/lexer"
)

func readProject(path string) (entries []fs.DirEntry) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func main() {
	fmt.Println("***FIX language compiler***")

	projectPath := os.Args[1]
	fmt.Println(projectPath)
	files := readProject(projectPath)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(projectPath, file.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		tokens := lexer.Tokenize(string(content))

        for _, token := range tokens {
            fmt.Println(token)
        }
	}
}

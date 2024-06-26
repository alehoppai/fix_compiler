package cli

import (
    "fmt"
    "os"
    "log"
    "io/fs"
	"path/filepath"
    "github.com/fix_compiler/lexer"
    "github.com/fix_compiler/parser"
)

const CLI_LOGO string = `
╔═════════════════════════╗    FIX language compiler.
║·························║┐   Developed by Aleh Belski <aleh_belski@outlook.com>
║···█████·██░·██░···██░···║│┐  pre-0.0.1v
║···██░░··██░··██░·██░····║││░
║···████··██░···███░······║││░
║···██░···██░··██░·██░····║││░
║···██░···██░·██░···██░···║││░
║····░░····░░··░░····░░···║││░
║·························║││░
╚═════════════════════════╝││░
 └─────────────────────────┘│░
  └─────────────────────────┘░
`

func readProject(path string) (entries []fs.DirEntry) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func performCompilation(projectPath string, debug bool) {
	files := readProject(projectPath)

	for i, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(projectPath, file.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

        fmt.Printf("Tokenizing %s -- %d/%d files.\n", filePath, i+1, len(files))
        if file.Name() == "structs.fix" {
            // TODO: struct parser not ready
            continue
        }
        tokens := lexer.Tokenize(string(content))
        if debug {
            PrettyPrintTokens(&tokens)
        }

        ast := parser.PrepareAST(tokens)
        if debug {
            for _, node := range ast {
                PrettyPrintAST(node, 0)
            }
        }
	}
}

func Start() {
    fmt.Print(CLI_LOGO)

    switch os.Args[1] {
        case "src":
            debug := len(os.Args) == 4 && os.Args[3] == "debug"
            performCompilation(os.Args[2], debug)
        case "help":
            fmt.Println("Help")

        default:
            fmt.Printf("Unknown argument %s", os.Args[1])
    }
}

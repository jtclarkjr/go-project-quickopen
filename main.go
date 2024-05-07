package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <editor> <path_to_project>")
		os.Exit(1)
	}

	editor := os.Args[1]
	projectPath := os.Args[2]
	var err error

	switch editor {
	case "xcode":
		err = openProject("Xcode", projectPath)
	case "vscode":
		err = openProject("Visual Studio Code", projectPath)
	// Add other IDE here
	default:
		fmt.Println("Unsupported editor. Use 'xcode' or 'vscode'.")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Failed to open project: %s\n", err)
		os.Exit(2)
	}
}

func openProject(editorName, path string) error {
	cmd := exec.Command("open", "-a", editorName, path)
	return cmd.Run()
}

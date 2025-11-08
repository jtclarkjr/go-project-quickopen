package main

import (
	"fmt"
	"os"

	"github.com/jtclarkjr/go-project-quickopen"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	args := os.Args[1:]

	if len(args) == 1 {
		switch args[0] {
		case "--version", "-v":
			printVersion()
			return nil
		case "--help", "-h":
			printHelp()
			return nil
		case "list":
			listEditors()
			return nil
		}
	}

	if len(args) > 2 {
		return fmt.Errorf("too many arguments. Usage: quickopen [editor] [path_to_project]")
	}

	editor, path := parseOpenArgs(args)

	if !quickopen.IsValidEditor(editor) {
		return fmt.Errorf("invalid IDE '%s'. Use 'quickopen list' to see available editors", editor)
	}

	return quickopen.OpenProject(editor, path)
}

func parseOpenArgs(args []string) (editor, path string) {
	switch len(args) {
	case 0:
		return "vscode", "."
	case 1:
		if quickopen.IsValidEditor(args[0]) {
			return args[0], "."
		}
		return "vscode", args[0]
	default:
		return args[0], args[1]
	}
}

func printVersion() {
	fmt.Printf("quickopen %s\n", quickopen.Version)
}

func printHelp() {
	fmt.Printf("quickopen %s\n\n", quickopen.Version)
	fmt.Println("USAGE:")
	fmt.Println("    quickopen [OPTIONS] [EDITOR] [PATH]")
	fmt.Println()
	fmt.Println("OPTIONS:")
	fmt.Println("    -h, --help       Print help information")
	fmt.Println("    -v, --version    Print version information")
	fmt.Println("    list             List available editors")
	fmt.Println()
	fmt.Println("ARGS:")
	fmt.Println("    [EDITOR]         Editor to use (default: vscode)")
	fmt.Println("    [PATH]           Path to open (default: current directory)")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("    quickopen                    # Open current directory in VS Code")
	fmt.Println("    quickopen cursor             # Open current directory in Cursor")
	fmt.Println("    quickopen /path/to/project   # Open path in VS Code")
	fmt.Println("    quickopen cursor ~/project   # Open path in Cursor")
}

func listEditors() {
	fmt.Println("Available editors:")
	editorKeys := quickopen.GetEditorKeys()
	for _, key := range editorKeys {
		fmt.Printf("  - %-12s (%s)\n", key, quickopen.Editors[key])
	}
}

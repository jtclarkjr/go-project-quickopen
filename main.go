package main

import (
	"fmt"
	"os"
	"os/exec"
)

const version = "1.0.0"

var editors = map[string]string{
	"xcode":      "Xcode",
	"vscode":     "Visual Studio Code",
	"webstorm":   "WebStorm",
	"pycharm-ce": "PyCharm CE",
	"intellij":   "IntelliJ IDEA",
	"goland":     "GoLand",
	"cursor":      "Cursor",
}

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

	editorName, ok := editors[editor]
	if !ok {
		return fmt.Errorf("invalid IDE '%s'. Use 'quickopen list' to see available editors", editor)
	}

	return openProject(editorName, path)
}

func parseOpenArgs(args []string) (editor, path string) {
	switch len(args) {
	case 0:
		return "vscode", "."
	case 1:
		if isValidEditor(args[0]) {
			return args[0], "."
		}
		return "vscode", args[0]
	default:
		return args[0], args[1]
	}
}

func isValidEditor(name string) bool {
	_, ok := editors[name]
	return ok
}

func printVersion() {
	fmt.Printf("quickopen %s\n", version)
}

func printHelp() {
	fmt.Printf("quickopen %s\n\n", version)
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
	editorKeys := []string{"xcode", "vscode", "webstorm", "pycharm-ce", "intellij", "goland", "cursor"}
	for _, key := range editorKeys {
		fmt.Printf("  - %-12s (%s)\n", key, editors[key])
	}
}

func openProject(editorName, path string) error {
	cmd := exec.Command("open", "-a", editorName, path)
	return cmd.Run()
}

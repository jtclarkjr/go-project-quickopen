package quickopen

import (
	"fmt"
	"os/exec"
)

// OpenProject opens a project with the specified editor and path
func OpenProject(editor, projectPath string) error {
	switch editor {
	case "xcode":
		return openProject("Xcode", projectPath)
	case "vscode":
		return openProject("Visual Studio Code", projectPath)
	// Add other IDE here
	default:
		return fmt.Errorf("unsupported editor. Use 'xcode' or 'vscode'")
	}
}

func openProject(editorName, path string) error {
	cmd := exec.Command("open", "-a", editorName, path)
	return cmd.Run()
}

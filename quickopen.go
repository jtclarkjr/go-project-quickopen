package quickopen

import (
	"fmt"
	"os/exec"
)

const Version = "1.0.0"

var Editors = map[string]string{
	"xcode":      "Xcode",
	"vscode":     "Visual Studio Code",
	"webstorm":   "WebStorm",
	"pycharm-ce": "PyCharm CE",
	"intellij":   "IntelliJ IDEA",
	"goland":     "GoLand",
	"cursor":     "Cursor",
}

func IsValidEditor(name string) bool {
	_, ok := Editors[name]
	return ok
}

func OpenProject(editor, path string) error {
	editorName, ok := Editors[editor]
	if !ok {
		return fmt.Errorf("invalid IDE '%s'", editor)
	}

	cmd := exec.Command("open", "-a", editorName, path)
	return cmd.Run()
}

func GetEditorKeys() []string {
	return []string{"xcode", "vscode", "webstorm", "pycharm-ce", "intellij", "goland", "cursor"}
}

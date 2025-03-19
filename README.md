# Project Quick Open

Quick open for projects in various IDEs, including Xcode, Visual Studio Code, WebStorm, IntelliJ IDEA, GoLand, Cursor, and Neovim. You can also add your own IDE to the package and script.

## Installation

If needed:

```bash
go mod tidy
```

## Usage

### Package

`/quickopen`

```go
package main

import (
    "fmt"
    quickopen "github.com/jtclarkjr/go-project-quickopen/quickopen"
)

func main() {
    err := quickopen.OpenProject("vscode", "/path/to/your/project")
    if err != nil {
        fmt.Printf("Failed to open project: %s\n", err)
    }
}
```

### Script

`root`

Two ways to run:

1. Move the `main` executable file to your user directory (`~/`) or projects directory so you don't have to navigate into this folder to run quick open.

Example commands:

```bash
./main xcode /path/to/your/project
./main vscode /path/to/your/project
./main webstorm /path/to/your/project
./main intellij /path/to/your/project
./main goland /path/to/your/project
./main cursor /path/to/your/project
./main neovim /path/to/your/project
```

---

2. Move the `main.go` file to your user directory (`~/`) or projects directory so you don't have to navigate into this folder to run quick open.

Example commands:

```bash
go run . xcode /path/to/your/project
go run . vscode /path/to/your/project
go run . webstorm /path/to/your/project
go run . intellij /path/to/your/project
go run . goland /path/to/your/project
go run . cursor /path/to/your/project
go run . neovim /path/to/your/project
```

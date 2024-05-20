# Project Quick open

Quick open for projects in xcode and vscode. Can also add your own IDE to package and script

## Installation

If needed:

```bash
go mod tidy
```

## Usages

### Package

`/quickopen`

```
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

Two ways to run.

1. move `main` executable file to user directory (~/) or projects directory so don't have to go into this folder to run quick open

xcode:

```bash
./main xcode /path/to/your/project
```

vscode:

```bash
./main vscode /path/to/your/project
```

---

2. move `main.go` file to user directory (~/) or projects directory so don't have to go into this folder to run quick open

xcode:

```bash
go run . xcode /path/to/your/project
```

vscode:

```bash
go run . vscode /path/to/your/project
```

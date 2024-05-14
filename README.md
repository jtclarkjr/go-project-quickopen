# Project Quick open

Quick open for projects in xcode and vscode. Can also add your own IDE to script

## Installation

If needed:

```bash
go mod tidy
```

## Usage

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

2. move `main.go` and `go.mod` files to user directory (~/) or projects directory so don't have to go into this folder to run quick open

xcode:

```bash
go run . xcode /path/to/your/project
```

vscode:

```bash
go run . vscode /path/to/your/project
```

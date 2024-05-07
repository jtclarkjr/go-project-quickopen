# Project Quick open

Quick open for projects in xcode and vscode. Can also add your own IDE to script

## Installation

If needed:

```bash
go mod tidy
```

## Usage

Two ways to run.

1. move `main` executable file to user directory (~/) main projects directory so don't have to go into this folder to run quick open

xcode:

```bash
./main xcode /path/to/your/project.xcodeproj
```

vscode:

```bash
./main vscode /path/to/your/projectFolder/
```

note: can do `./` instead of `main.go` also

---

2. move `main.go` and `go.mod` files to user directory (~/) or main projects directory so don't have to go into this folder to run quick open

xcode:

```bash
go run main.go xcode /path/to/your/project.xcodeproj
```

vscode:

```bash
go run main.go vscode /path/to/your/projectFolder/
```

note: can do instead of `main.go` also

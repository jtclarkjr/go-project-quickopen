# Project Quick Open

Quick open for projects in various IDEs, including Xcode, Visual Studio Code, WebStorm, IntelliJ IDEA, GoLand, and Cursor.

## Installation

1. Build the binary:

```bash
go build -o quickopen .
```

2. Install globally:

```bash
sudo cp quickopen /usr/local/bin/
```

## Usage

```bash
quickopen [OPTIONS] [EDITOR] [PATH]
```

### Options

- `-h, --help` - Print help information
- `-v, --version` - Print version information
- `list` - List available editors

### Arguments

- `[EDITOR]` - Editor to use (default: vscode)
- `[PATH]` - Path to open (default: current directory)

### Examples

```bash
quickopen                    # Open current directory in VS Code
quickopen cursor             # Open current directory in Cursor
quickopen /path/to/project   # Open path in VS Code
quickopen cursor ~/project   # Open path in Cursor
quickopen list               # List available editors
```

### Available Editors

- `xcode` - Xcode
- `vscode` - Visual Studio Code
- `webstorm` - WebStorm
- `pycharm-ce` - PyCharm CE
- `intellij` - IntelliJ IDEA
- `goland` - GoLand
- `cursor` - Cursor

# qd - Quick Directory

`qd` is a lightweight and straightforward CLI tool to manage directory aliases. With `qd`, you can save frequently used directories as aliases and quickly navigate to them from your terminal.

## Installation

Clone the repository:
```bash
git clone https://github.com/enisbt/qd.git
cd qd
```
Build the binary:
```bash
go build -o qd
```
Move the binary to a directory in your $PATH:
```bash
sudo mv qd /usr/local/bin
```
Add the shell function for persistent navigation. Add the following to your ~/.bashrc or ~/.zshrc:
```bash
function qd() {
    if [[ $1 == "save" ]]; then
        /usr/local/bin/qd save "$2"
    elif [[ $1 == "list" ]]; then
        /usr/local/bin/qd list
    elif [[ $1 == "delete" ]]; then
        /usr/local/bin/qd delete "$2"
    else
        cd "$(/usr/local/bin/qd "$1")"
    fi
}
```
Reload your shell configuration
```bash
source ~/.bashrc  # or `source ~/.zshrc`
```

## Usage

Save current directory as an alias

```bash
qd save <alias>
```

Navigate to a saved alias

```bash
qd <alias>
```

List all aliases
```bash
qd list
```

Delete alias
```bash
qd delete <alias>
```

## License

Distributed under the MIT License. See `LICENSE` for more information.
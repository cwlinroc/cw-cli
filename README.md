# cw-cli

A personal CLI toolkit for scaffolding files, converting encodings, and saving screenshots.

## Installation

```sh
go build -o cw
```

## Development

```sh
golangci-lint run
go test ./...
go build ./...
```

## Commands

| Command                     | Description                                                        |
| --------------------------- | ------------------------------------------------------------------ |
| [`add`](docs/add.md)        | Interactively pick a directory, then scaffold a file               |
| [`touch`](docs/touch.md)    | Scaffold a file in the current working directory                   |
| [`bsf`](docs/bsf.md)        | Convert between Unicode escape sequences and UTF-8 characters      |
| [`lf` / `crlf`](docs/lf.md) | Convert line endings on a file or an entire directory tree         |
| [`utf8`](docs/utf8.md)      | Convert files to UTF-8, stripping BOM or re-encoding from Big5     |
| [`ss`](docs/ss.md)          | Save a clipboard image to a PNG or BMP file _(Windows/macOS only)_ |
| [`path`](docs/path.md)      | Copy the current directory path to the clipboard _(Windows/macOS only)_ |
| [`timer`](docs/timer.md)    | Display a live countdown timer in the terminal                     |

## Quick start

```sh
# scaffold a C# class in a directory you pick interactively
cw add cs

# scaffold a Razor page in the current directory
cw touch razor Index

# scaffold a customized .gitignore for .NET and Go
cw touch gitignore --net --go

# convert all source files to LF line endings
cw lf -a

# convert a file to UTF-8
cw utf8 legacy.txt

# save clipboard screenshot
cw ss

# start a 1 hour 30 minute countdown timer
cw timer 1h30m
```

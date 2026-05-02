# AGENTS.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```
go build -o cw          # build binary
go build ./...          # build all packages
go run .                # run from source
go mod tidy             # sync module metadata after dependency changes
```

## Tests & Lint

```
go test ./...                                           # all tests
go test ./internal/template                             # one package
go test ./internal/template -run '^TestGenerateCs$' -count=1  # one test
go vet ./...                                            # vet
gofmt -w main.go cmd/*.go internal/**/*.go              # format
```

No test files exist yet; `go test ./...` passes because there is nothing to fail.

## Architecture

```
main.go                       entry point → cmd.Execute()
cmd/
  root.go                     rootCmd, huhTheme (huh.ThemeBase16())
  add.go                      interactive flow: prompts for dir, then name
  touch.go                    quick creation: uses cwd, optional name arg
  bsf.go  lf.go  utf8.go     encoding/line-ending utilities
  ss.go                       screenshot save (//go:build !linux only)
internal/
  file/
    direct.go                 PickDir() — interactive Huh directory picker
    name.go                   GetName() — interactive or arg-based name input
    dotnet_check.go           ExtractCSNamespace() — walks up to nearest .csproj
  template/
    cs.go                     GenerateCs(targetDir, fileName)
    razor.go                  GenerateRazor(targetDir, fileName)
    code.go                   GenerateCode(targetDir, fileName)
    page.go                   GeneratePage(targetDir)
    editorConfig.go           GenerateEditorConfig(targetDir)
```

**`add` vs `touch`:** `add <type>` is interactive (calls `PickDir` + `GetName` via Huh prompts). `touch <type> [name]` is non-interactive (uses `os.Getwd()` + optional name arg).

**Namespace detection:** `ExtractCSNamespace` in `internal/file/dotnet_check.go` walks up from `targetDir` up to 50 levels, collecting directory names, until it finds a `.csproj`. The collected names are reversed and joined with `.` to form the namespace.

**`ss` command:** Excluded from Linux via `//go:build !linux`. BMP output is actually PNG data — the current code saves PNG bytes to a `.bmp` file by design (see `saveAsBmp` in `cmd/ss.go`).

**Global theme:** `huhTheme` in `cmd/root.go` (not `hunTheme`).

## Smoke Tests

```
go run . touch cs MyClass
go run . touch razor Index
go run . touch code app
go run . touch page
go run . touch editorconfig
go run . add cs
go run . ss                   # Windows/macOS only
```

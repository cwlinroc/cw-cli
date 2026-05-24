# AGENTS.md

This file guides agentic coding tools working in `C:\Users\cwlin\projects\my-initial-setup\cw-cli`.

## Environment

- Module path: `cw`
- Go version: `1.25`
- Entry point: `main.go`
- Command package: `cmd`
- Reusable helpers: `internal/`

## Layout

```text
main.go                       Calls cmd.Execute()
cmd/
  root.go                     Root Cobra command and global Huh theme
  add.go                      Interactive creation commands
  touch.go                    Non-interactive creation commands
  clean.go                    Recursive cleanup command
  bsf.go                      Unicode escape converter
  lf.go                       LF/CRLF conversion utilities
  utf8.go                     UTF-8 conversion utilities
  ss.go                       Clipboard screenshot save command (!linux only)
  path.go                     Clipboard path copy command (!linux only)
  timer.go                    Countdown timer with live TUI (bubbletea + bubbles/timer)
internal/
  file/
    direct.go                 Interactive directory picker
    name.go                   Name input helpers
    dotnet_check.go           C# namespace discovery via nearest .csproj
  template/
    cs.go                     C# class generation
    razor.go                  Razor page generation
    code.go                   VS Code workspace generation
    page.go                   Svelte +page.svelte generation
    editorConfig.go           .editorconfig generation
    gitignore.go              .gitignore file generation
    gitignore_test.go         Unit tests for .gitignore generation
docs/
  *.md                        Command docs
```

## Build

Run from the repository root.

```bash
go build -o cw
go build ./...
go run .
go run . --help
go mod tidy
```

- Use `go build ./...` as the broad compile check.
- Run `go mod tidy` only when dependencies changed.

## Test, Lint, Format

```bash
go test ./...
go vet ./...
golangci-lint run
gofmt -w main.go cmd/*.go internal/**/*.go
```

Current state:

- There are now several `_test.go` files under `internal/template/`.
- `go test ./...` runs all these template verification tests.

When tests are added or executed, use these patterns:

```bash
go test ./...
go test ./cmd
go test ./internal/template
go test ./internal/template -run '^TestGenerateCs$' -count=1
go test ./internal/template -run '^TestGenerateGitignore$' -count=1
go test ./cmd -run '^TestTouchCsCmdRun$' -count=1
```

- Prefer package-scoped runs while iterating.
- Use `-run` with an anchored regex for a single test.
- Use `-count=1` to avoid cached results while debugging.
- Run `golangci-lint run` before finishing Go changes.

## Fast Smoke Tests

```bash
go run . touch cs MyClass
go run . touch razor Index
go run . touch code app
go run . touch page
go run . touch editorconfig
go run . touch gitignore --net --go
go run . add cs
go run . add gitignore
go run . clean
go run . utf8 --help
go run . lf --help
go run . crlf --help
go run . ss
go run . timer 5s
go run . timer 1hr30min
```

- `cmd/ss.go` is guarded by `//go:build !linux`.
- `cw timer` supports units: s/sec, m/min, h/hr, and compounds like "1h30m".

## Architecture Notes

- `main.go` should stay thin and only delegate to `cmd.Execute()`.
- Commands are registered in `init()` functions inside `cmd/*.go`.
- `cmd/` owns CLI UX, flags, prompts, and user-facing output.
- `internal/file` contains prompt and path helpers.
- `internal/template` contains file generation logic and template text.
- `add <type>` is interactive and prompts for a directory and sometimes a name.
- `touch <type> [name]` is non-interactive and uses the current working directory.

Repository-specific behavior:

- `ExtractCSNamespace` walks upward until it finds a `.csproj`, then builds a dotted namespace from path segments.
- `GenerateRazor` creates both `.cshtml` and `.cshtml.cs` files.
- `GeneratePage` writes `+page.svelte` based on the folder name.
- `saveAsBmp` currently writes PNG-encoded bytes to a `.bmp` filename by design in this repo.

## Code Style

Follow idiomatic Go first, then preserve local consistency in touched files.

### Formatting And Imports

- Always run `gofmt` on modified Go files.
- Use standard Go import grouping: stdlib first, then third-party imports.
- Do not keep unused imports.

### Naming

- Exported names use PascalCase.
- Unexported names use camelCase.
- Prefer idiomatic new names like `targetDir`, `namespace`, `baseDir`, and `fileName`.
- Some existing names are non-idiomatic, such as `targetDirect`, `nameSpace`, and `lf_to_crlf`; do not rename them unless the change is tightly scoped.

### Types And Functions

- Prefer concrete types unless an interface is needed for behavior or testing.
- Keep helpers near their usage when they are file-local.
- Avoid introducing new packages for small refactors.
- Keep control flow simple and prefer early returns.

### Error Handling

- In `cmd/`, errors are usually printed with `fmt.Println` or `fmt.Printf` and the command returns.
- In `internal/`, prefer returning errors upward rather than printing there.
- For new code, prefer `fmt.Errorf("...: %w", err)` when wrapping errors.
- Avoid panics for normal CLI failures.

### File And CLI Behavior

- Use `os.ReadFile` and `os.WriteFile` for simple whole-file operations.
- Use `filepath.Join` for filesystem paths.
- Close created files and handle close errors when they matter to the command result.
- Check write errors in new file-generation code.
- Add new commands with `rootCmd.AddCommand(...)` in `init()`.
- Prefer Cobra validators such as `cobra.NoArgs`, `cobra.ExactArgs`, and `cobra.RangeArgs`.
- Reuse the global `huhTheme` for interactive prompts.

## Agent Guidance

- Make minimal, targeted changes.
- Preserve current behavior unless the task explicitly requires changing it.
- If you add tests, place them next to the package they cover.
- If you change templates or command UX, run at least one relevant `go run . ...` smoke test.
- If you change shared helpers such as `getAllProgramFiles`, `ExtractCSNamespace`, or line-ending conversion logic, run `go test ./...` and at least one manual command path.

## External Rules

- No `.cursorrules` file found.
- No `.cursor/rules/` directory found.
- No `.github/copilot-instructions.md` file found.

If any of those files are added later, their instructions should be treated as additional repository policy.

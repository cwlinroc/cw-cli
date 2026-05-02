# AGENTS.md
This file is for coding agents working in `cw`.
It replaces the older placeholder-style project notes with repo-specific guidance.

## Scope
- Repository: `cw`
- Language: Go
- Module path: `cw`
- Main artifact: CLI binary named `cw`
- Entry point: `main.go` -> `cmd.Execute()`

## External Instruction Sources
- Included: `.github/copilot-instructions.md`
- Included here by summary, not by reference only
- Not found: `.cursorrules`
- Not found: `.cursor/rules/`

## Repo Summary
- `cmd/` contains Cobra commands and CLI-facing behavior.
- `template/` contains file generation logic and embedded templates.
- Main flows: `cw add <type>` for interactive creation and `cw touch <type> [name]` for quick creation.
- Supported generators: C# classes, Razor pages, VS Code workspaces, Svelte pages, and `.editorconfig` files.
- Utility commands include encoding conversion, line-ending conversion, unicode conversion, screenshot saving on non-Linux builds, and cleanup.

## Architecture Notes
- Follow the Cobra pattern already used here: define command vars at package scope and register them in `init()`.
- Keep command wiring in `cmd/` and file/template generation in `template/`.
- The global Huh theme is `hunTheme = huh.ThemeBase16()` in `cmd/root.go`.
- Namespace detection for generated C# files lives in `template/cs.go` and walks upward until it finds a `.csproj`.
- Template functions currently follow `Generate<Type>(targetDir, fileName) error` or `Generate<Type>(targetDir) error`.
- There is no external runtime config system in this repo today.

## Build Commands
- Build the CLI binary: `go build -o cw`
- Build all packages: `go build ./...`
- Run from source: `go run .`
- Run built binary: `./cw`
- Refresh module metadata when dependencies change: `go mod tidy`

## Test Commands
- Run all tests: `go test ./...`
- Run one package: `go test ./template`
- Run a single test by exact name: `go test ./template -run '^TestGenerateCs$' -count=1`
- Run a single subtest: `go test ./template -run '^TestGenerateCs$/^creates file$' -count=1`
- Run tests verbosely: `go test ./template -v`
- Disable test caching while iterating: add `-count=1`

## Current Test State
- At the time this file was written, the repo has no `*_test.go` files.
- The single-test commands above are the correct Go patterns to use once tests are added.
- `go test ./template` succeeds today because that package has no CGO dependency.

## Lint And Formatting Commands
- Format Go code: `gofmt -w main.go cmd/*.go template/*.go`
- Vet packages: `go vet ./...`
- There is no repo-local `golangci-lint` or `staticcheck` config at present.
- If you use `goimports`, keep the result stable and do not introduce import churn unrelated to the task.

## Verified Environment Caveat
- `cmd/ss.go` is excluded from Linux builds with a build tag because `golang.design/x/clipboard` requires X11 headers such as `X11/Xlib.h`.
- On Linux, the `ss` command is intentionally unavailable.
- In this Linux environment, `go test ./...` now succeeds because the screenshot command is not compiled.
- On non-Linux builds, the screenshot command still depends on clipboard support from `golang.design/x/clipboard`.

## Useful Manual Smoke Tests
- Quick C# class generation: `go run . touch cs MyClass`
- Quick Razor generation: `go run . touch razor Index`
- Quick workspace generation: `go run . touch code app`
- Quick Svelte page generation: `go run . touch page`
- Quick editorconfig generation: `go run . touch editorconfig`
- Interactive flow: `go run . add cs`
- Screenshot save smoke test on Windows or other non-Linux builds: `go run . ss`

## Imports
- Let `gofmt` be the minimum source of truth for formatting.
- Prefer standard-library imports first, then third-party imports, then local module imports if your formatter supports grouping.
- Do not hand-tune import order unless the formatter requires it.
- Remove unused imports rather than leaving placeholders.

## Formatting
- Use normal Go formatting and let `gofmt` decide spacing and indentation.
- Keep functions and files compact when the logic is small.
- Prefer early returns over deep nesting.
- Avoid adding blank lines unless they separate distinct steps.
- Do not reformat unrelated files just because you touched the repo.

## Types And Signatures
- Prefer concrete types unless an interface is already required.
- Keep function signatures small and purpose-specific.
- Return `error` for operations that can fail; do not hide failures.
- Favor helpers that accept explicit paths and names over hidden global state.
- Use `filepath` for filesystem paths.

## Naming Conventions
- For new Go identifiers, prefer idiomatic Go mixedCaps names.
- Exported names should be rare and justified by package boundaries.
- Unexported helpers are preferred unless another package needs the symbol.
- Command vars should stay descriptive and consistent with current patterns, such as `touchCsCmd` and `touchCsCmdRun`.
- Template entry points should keep the existing `Generate<Type>` naming pattern.
- Some older private helpers use snake_case, such as `cs_namespace` and `lf_to_crlf`.
- Do not rename existing symbols only for style unless the task is explicitly a refactor.

## Error Handling
- In library-style helpers, return errors upward instead of printing and swallowing them.
- In Cobra `Run` functions, printing a user-facing error and returning early matches the current repo style.
- When adding new wrapped errors, prefer `fmt.Errorf("context: %w", err)`.
- Avoid `panic`, `log.Fatal`, and silent failure paths in normal command execution.
- Do not print the same error at multiple layers unless that duplication is intentional for UX.

## File And IO Conventions
- Use `os.ReadFile`, `os.WriteFile`, `os.Create`, and `os.ReadDir` as the repo already does.
- Use `0644` for regular files and `0755` for created directories unless there is a strong reason otherwise.
- Join paths with `filepath.Join()`.
- Preserve existing behavior around generated filenames and extensions.
- Avoid broad destructive operations outside commands explicitly meant for cleanup.

## Command Design Guidelines
- Register every command with `rootCmd.AddCommand()` from an `init()` function.
- Parse flags via Cobra APIs such as `cmd.Flags().GetBool()`.
- Keep argument validation in Cobra metadata when possible, such as `cobra.NoArgs`, `cobra.ExactArgs`, or `cobra.RangeArgs`.
- Put business logic in small helpers when that makes the code easier to test.
- Keep interactive Huh prompt usage inside CLI-facing code, not inside reusable template helpers.

## Template Guidelines
- Keep templates embedded as string literals unless there is a clear need for external files.
- Keep template generation deterministic.
- When creating multiple files for one command, fail clearly if any file creation step fails.
- If a generator depends on repo structure, make that assumption obvious in the error path.
- Preserve the current C# namespace detection behavior unless intentionally changing it.

## Testing Guidance For New Code
- Prefer table-driven tests for pure helpers.
- Use `t.TempDir()` for file-generation tests.
- Assert on both file existence and file contents.
- Keep interactive prompt code thin so core behavior can be tested without terminal UI.
- If you add tests under `cmd/`, remember that `ss.go` only builds on non-Linux platforms.
- Keep clipboard-dependent coverage isolated behind build constraints rather than weakening the rest of the package tests.

## Repo-Specific Gotchas
- The root command metadata in `cmd/root.go` is still mostly Cobra scaffold text.
- Some code mixes printing and returning errors; prefer reducing that in new changes rather than expanding it.
- `template/` is the safest package for isolated unit tests in the current environment.
- The screenshot command currently writes PNG data even for `.bmp` output; preserve or change that only intentionally.
- The screenshot command is intentionally unavailable on Linux.
- The repo root `.editorconfig` asks for UTF-8, trailing newline, trimmed trailing whitespace, and no markdown trailing-whitespace trimming.

## Agent Expectations
- Make the smallest correct change.
- Preserve the interactive-vs-quick command split described in the Copilot instructions.
- Keep command registration, template delegation, and namespace logic easy to find.
- Verify with the narrowest useful command first, then broaden if the environment supports it.
- Remember that Linux and non-Linux builds differ for the `ss` command because of the build constraint.

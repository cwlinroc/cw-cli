# touch

Non-interactive version of [`add`](add.md). Creates a file in the **current working directory** without any prompts. An optional name argument can be passed; if omitted, you are prompted for one.

## Subcommands

### `touch cs [name]`

```sh
cw touch cs MyService
# → writes MyService.cs in cwd
```

Namespace is derived by walking up from cwd to the nearest `.csproj`.

---

### `touch razor [name]`

```sh
cw touch razor Login
# → writes Login.cshtml + Login.cshtml.cs in cwd
```

---

### `touch code [name]`

```sh
cw touch code MyWorkspace
# → writes MyWorkspace.code-workspace in cwd
```

---

### `touch page`

```sh
cw touch page
# → writes +page.svelte in cwd
```

---

### `touch editorconfig`

```sh
cw touch editorconfig
# → writes .editorconfig in cwd
```

---

### `touch gitignore`

```sh
cw touch gitignore [flags]
# → writes .gitignore in cwd
```

Supports selective ecosystem ignore rules via flags. If no flags are provided, all supported languages are included.

**Supported languages & flags:**
- `.NET`: `--net`, `--dotnet`
- `Go`: `--go`, `--golang`
- `Java`: `--java`
- `JS/TS` (npm/pnpm/yarn): `--js`, `--ts`, `--node`, `--npm`, `--pnpm`
- `Python`: `--py`, `--python`

Example:
```sh
cw touch gitignore --go --py
# → writes a .gitignore only including Go and Python ignore patterns
```

---

## Comparison: add vs touch

|           | `add`                                 | `touch`                          |
| --------- | ------------------------------------- | -------------------------------- |
| Directory | Interactive picker                    | Current working directory        |
| Use case  | Placing files anywhere in the project | Quick creation from the terminal |


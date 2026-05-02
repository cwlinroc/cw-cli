# lf / crlf

Convert line endings in source files. `lf` normalises to Unix line endings (`\n`); `crlf` normalises to Windows line endings (`\r\n`).

Multi-byte UTF-8 sequences are never corrupted — the converter is byte-aware and only touches `\r` and `\n` bytes.

## Usage

```
cw lf   [filePath]  [-a]
cw crlf [filePath]  [-a]
```

## Flags

| Flag          | Description                                                            |
| ------------- | ---------------------------------------------------------------------- |
| `-a`, `--all` | Recursively convert all recognised text files in the current directory |

## Examples

```sh
# convert a single file
cw lf  src/main.cs
cw crlf src/main.cs

# convert all source files under cwd
cw lf -a
cw crlf -a
```

## Supported file extensions

`.py` `.cs` `.go` `.java` `.js` `.ts` `.cshtml` `.razor` `.csproj` `.sln` `.html` `.css` `.json` `.xml` `.yaml` `.yml` `.md` `.txt` `.sh` `.bat` `.cmd` `.gitignore` `.dockerignore` `.editorconfig` `.vue` `.php` `.rb` `.cpp` `.h` `.hpp` `.jsx` `.tsx` `.sql` `.swift` `.svelte` `.mod` `.sum`

## Skipped directories

`.git` `.vscode` `node_modules` `bin` `obj` `packages` `dist` `build` `out` `target` `__pycache__` `.idea` `.vs` `.nuxt`


# utf8

Convert files to UTF-8 encoding. Handles BOM removal and re-encoding from Big5 (Traditional Chinese).

## Usage

```
cw utf8 [filePath] [-a] [-b]
```

## Flags

| Flag                 | Description                                                            |
| -------------------- | ---------------------------------------------------------------------- |
| `-a`, `--all`        | Recursively convert all recognised text files in the current directory |
| `-b`, `--accept-bom` | Keep UTF-8 BOM instead of stripping it                                 |

## Examples

```sh
# convert a single file
cw utf8 legacy.txt

# convert all source files under cwd
cw utf8 -a

# convert but preserve BOM
cw utf8 -b legacy.txt
```

## Conversion logic

1. **BOM detected** (`EF BB BF`) — strips it unless `-b` is set.
2. **Already valid UTF-8** — no change.
3. **Big5 encoded** — decoded and re-written as UTF-8.

Additional encodings can be added by extending the source.


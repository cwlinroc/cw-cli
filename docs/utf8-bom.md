# utf8-bom

Convert files to UTF-8 with BOM encoding. Re-encodes Big5 (Traditional Chinese) input and ensures UTF-8 files contain exactly one BOM.

## Usage

```
cw utf8-bom [filePath] [-a]
```

## Flags

| Flag          | Description                                                                      |
| ------------- | -------------------------------------------------------------------------------- |
| `-a`, `--all` | Recursively convert all recognised text files in the current directory to UTF-8 with BOM |

## Examples

```sh
# convert a single file
cw utf8-bom legacy.txt

# convert all source files under cwd
cw utf8-bom -a
```

## Conversion logic

1. **BOM detected** (`EF BB BF`) - leaves the file unchanged.
2. **Already valid UTF-8** - adds a BOM.
3. **Big5 encoded** - decodes and re-writes the file as UTF-8 with BOM.

Additional encodings can be added by extending the source.

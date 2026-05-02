# ss

Save the current clipboard image to a file. Supports PNG and BMP output.

> **Platform:** Windows and macOS only. Not available on Linux.

## Usage

```
cw ss [name]
```

## Arguments

| Argument     | Description                          |
| ------------ | ------------------------------------ |
| _(none)_     | Save as `YYYYMMDD_HHMMSS.png` in cwd |
| `png`        | Save as `YYYYMMDD_HHMMSS.png`        |
| `bmp`        | Save as `YYYYMMDD_HHMMSS.bmp`        |
| `<name>`     | Save as `<name>.png`                 |
| `<name>.png` | Save as `<name>.png`                 |
| `<name>.bmp` | Save as `<name>.bmp`                 |

## Examples

```sh
cw ss                 # → 20240501_143022.png
cw ss diagram         # → diagram.png
cw ss diagram.bmp     # → diagram.bmp
cw ss bmp             # → 20240501_143022.bmp
```

## Notes

- The image is read directly from the clipboard — copy an image first.
- BMP output is stored as PNG bytes with a `.bmp` extension (no native Go BMP encoder).


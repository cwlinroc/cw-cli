# bsf

Bidirectional converter between Unicode escape sequences (`\uXXXX`) and UTF-8 characters.

## Usage

```
cw bsf <value>
```

Requires exactly one argument.

## Conversion rules

| Input form                           | Output                                 |
| ------------------------------------ | -------------------------------------- |
| `uXXXX` (4 hex digits, no backslash) | UTF-8 character                        |
| `\uXXXX` (backslash-prefixed)        | UTF-8 character                        |
| Single character                     | `\uXXXX` escape                        |
| Multi-character string               | Unquoted / interpreted escape sequence |

## Examples

```sh
cw bsf u0041      # → A
cw bsf 'A'   # → A
cw bsf A          # → A
cw bsf u4e2d      # → 中
cw bsf 中         # → 中
```


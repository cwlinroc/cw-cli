# timer

Display a live countdown timer in the terminal.

## Usage

```
cw timer <duration>
```

## Description

The `timer` command shows a live, updating countdown timer using Bubble Tea. When the time is up, it displays a completion message and plays a terminal bell.

You can press `q` or `Ctrl+C` to abort the timer at any time. When the timer is finished, press any key to dismiss the message and exit.

## Duration Formats

The `<duration>` argument can be a simple number (treated as seconds) or a combination of numbers and units.

Supported units:
- **Seconds:** `s`, `sec`
- **Minutes:** `m`, `min`
- **Hours:** `h`, `hr`

### Examples

```sh
# 30 seconds (bare number)
cw timer 30

# 23 seconds
cw timer 23s
cw timer 23sec

# 13 minutes
cw timer 13m
cw timer 13min

# 5 hours
cw timer 5h
cw timer 5hr

# 1 hour and 30 minutes (compound duration)
cw timer 1h30m
cw timer 1hr30min
```

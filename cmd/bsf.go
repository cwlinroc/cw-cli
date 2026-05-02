package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(bsfCmd)
}

var bsfCmd = &cobra.Command{
	Use:   "bsf",
	Short: "convert utf-8 backslash format",
	Long:  `convert utf-8 backslash format. For example: cw bsf \uf240`,
	Args:  cobra.ExactArgs(1),
	Run:   bsfCmdRun,
}

//------------------------------------------------------------

func bsfCmdRun(cmd *cobra.Command, args []string) {
	chars := []rune(args[0])

	if chars[0] == 'u' && len(chars) == 5 {
		fmt.Println(fromEscaped(chars))
		return
	} else if chars[0] == '\\' && chars[1] == 'u' && len(chars) == 6 {
		fmt.Println(fromEscaped(chars[1:]))
		return
	} else if len(chars) == 1 {
		fmt.Println(toEscape(chars[0]))
		return
	} else {
		fmt.Println(format(args[0]))
		return
	}
}

func format(str string) string {
	quotedStr := fmt.Sprintf("\"%s\"", str)

	utf8Str, err := strconv.Unquote(quotedStr)
	if err != nil {
		return "Failed to unquote string"
	}

	return utf8Str
}

func toEscape(char rune) string {
	return fmt.Sprintf("\\u%04x", char)
}

func fromEscaped(chars []rune) string {
	target := chars[1:]

	if len(target) != 4 {
		return "Invalid hexadecimal character"
	}

	var bytes [3]byte

	fmt.Printf("%08b\n", bytes)

	for i, c := range target {
		var hex int
		switch {
		case c >= '0' && c <= '9':
			hex = int(c - '0')
		case c >= 'a' && c <= 'f':
			hex = int(c - 'a' + 10)
		case c >= 'A' && c <= 'F':
			hex = int(c - 'A' + 10)
		default:
			return "Invalid format"
		}

		fmt.Println(hex)

		switch i {
		case 0:
			bytes[0] = byte(hex | 0xe0)
		case 1:
			bytes[1] = byte(hex<<2 | 0x80)
		case 2:
			bytes[1] |= byte(hex >> 2)
			bytes[2] = byte(((hex & 0x03) << 4) | 0x80)
		case 3:
			bytes[2] |= byte(hex)
		}

		fmt.Printf("%08b\n", bytes)
	}

	return string(bytes[:])
}

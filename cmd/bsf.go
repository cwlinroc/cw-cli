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

	codePoint, err := strconv.ParseInt(string(target), 16, 32)
	if err != nil {
		return "Invalid format"
	}

	return string(rune(codePoint))
}

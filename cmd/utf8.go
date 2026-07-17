/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func init() {
	rootCmd.AddCommand(utf8Cmd)
	rootCmd.AddCommand(utf8BOMCmd)

	utf8Cmd.Flags().BoolP("all", "a", false, "Convert all files in the current directory to UTF-8")
	utf8Cmd.Flags().BoolP("accept-bom", "b", false, "Accept UTF-8 with BOM")
	utf8BOMCmd.Flags().BoolP("all", "a", false, "Convert all files in the current directory to UTF-8 with BOM")
}

var utf8Cmd = &cobra.Command{
	Use:   "utf8",
	Short: "Convert a file to UTF-8",
	Long:  `Convert a file to UTF-8. For example: cw utf8 file.txt`,
	Args:  cobra.RangeArgs(0, 1),
	Run:   utf8CmdRun,
}

var utf8BOMCmd = &cobra.Command{
	Use:   "utf8-bom",
	Short: "Convert a file to UTF-8 with BOM",
	Long:  `Convert a file to UTF-8 with BOM. For example: cw utf8-bom file.txt`,
	Args:  cobra.RangeArgs(0, 1),
	Run:   utf8BOMCmdRun,
}

type bomPolicy int

const (
	removeBOM bomPolicy = iota
	preserveBOM
	addBOM
)

var utf8BOM = []byte{0xEF, 0xBB, 0xBF}

func utf8CmdRun(cmd *cobra.Command, args []string) {
	acceptBOM, err := cmd.Flags().GetBool("accept-bom")
	if err != nil {
		fmt.Println("Error getting accept-bom flag: ", err)
		return
	}

	policy := removeBOM
	if acceptBOM {
		policy = preserveBOM
	}

	runUTF8Cmd(cmd, args, policy, "UTF-8")
}

func utf8BOMCmdRun(cmd *cobra.Command, args []string) {
	runUTF8Cmd(cmd, args, addBOM, "UTF-8 with BOM")
}

func runUTF8Cmd(cmd *cobra.Command, args []string, policy bomPolicy, encoding string) {
	all, err := cmd.Flags().GetBool("all")
	if err != nil {
		fmt.Println("Error getting all flag: ", err)
		return
	}

	if len(args) == 0 && !all {
		fmt.Println("Please provide a file path")
		return
	}

	if all {
		err = spinner.New().
			Title("Converting all files to " + encoding).
			Action(func() { convertAllFilesEncoding(policy, encoding) }).
			Run()
		if err != nil {
			fmt.Println("Error running spinner: ", err)
		}

		return
	}

	path := args[0]

	err = convertEncoding(path, true, policy)
	if err != nil {
		fmt.Println("Error converting encoding: ", err)
		return
	}
}

func convertAllFilesEncoding(policy bomPolicy, encoding string) {
	fmt.Println("Converting all files in the current directory to " + encoding)

	files, err := getAllProgramFiles()
	if err != nil {
		fmt.Println("Error getting all files: ", err)
		return
	}

	for _, file := range files {
		err := convertEncoding(file, false, policy)
		if err != nil {
			fmt.Println("Error converting encoding: ", err)
			return
		}
	}
}

func convertEncoding(path string, verbose bool, policy bomPolicy) error {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return err
	}

	utf8withBom := len(content) >= len(utf8BOM) &&
		content[0] == utf8BOM[0] &&
		content[1] == utf8BOM[1] &&
		content[2] == utf8BOM[2]

	if utf8withBom {
		if policy != removeBOM {
			return nil
		}

		content = content[len(utf8BOM):]
		err := os.WriteFile(path, content, 0o644)
		if err != nil {
			fmt.Println("Convert UTF-8 with BOM failed. Error writing file: ", err)
			return err
		}

		fmt.Println(path + " converted to UTF-8 from UTF-8 with BOM")

		return nil
	}

	if utf8.Valid(content) {
		if policy == addBOM {
			content = append(append([]byte{}, utf8BOM...), content...)
			if err := os.WriteFile(path, content, 0o644); err != nil {
				fmt.Println("Convert to UTF-8 with BOM failed. Error writing file: ", err)
				return err
			}

			fmt.Println(path + " converted to UTF-8 with BOM")
			return nil
		}

		if verbose {
			fmt.Println(path + " is UTF-8")
		}
		return nil
	}

	big5Encoder := traditionalchinese.Big5.NewDecoder()

	encoderPairs := []encoderPair{{big5Encoder, "Big5"}}

	for _, pair := range encoderPairs {
		exactEncode, err := convertToUTF8(path, content, pair.encoder, policy == addBOM)
		if err != nil {
			fmt.Println("Error converting to UTF-8: ", err)
			return err
		}

		if exactEncode {
			encoding := "UTF-8"
			if policy == addBOM {
				encoding = "UTF-8 with BOM"
			}
			fmt.Println(path + " converted to " + encoding + " from " + pair.encoding)
			return nil
		}
	}

	return nil
}

type encoderPair struct {
	encoder  transform.Transformer
	encoding string
}

func convertToUTF8(path string, content []byte, decoder transform.Transformer, withBOM bool) (bool, error) {
	utf8Content, _, err := transform.Bytes(decoder, content)
	if err != nil {
		return false, nil
	}
	if withBOM {
		utf8Content = append(append([]byte{}, utf8BOM...), utf8Content...)
	}

	err = os.WriteFile(path, utf8Content, 0o644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return true, err
	}

	return true, nil
}

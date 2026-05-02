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

	utf8Cmd.Flags().BoolP("all", "a", false, "Convert all files in the current directory to UTF-8")
	utf8Cmd.Flags().BoolP("accept-bom", "b", false, "Accept UTF-8 with BOM")
}

var utf8Cmd = &cobra.Command{
	Use:   "utf8",
	Short: "Convert a file to UTF-8",
	Long:  `Convert a file to UTF-8. For example: cw utf8 file.txt`,
	Args:  cobra.RangeArgs(0, 1),
	Run:   utf8CmdRun,
}

func utf8CmdRun(cmd *cobra.Command, args []string) {
	all, err := cmd.Flags().GetBool("all")
	if err != nil {
		fmt.Println("Error getting all flag: ", err)
		return
	}

	acceptBOM, err := cmd.Flags().GetBool("accept-bom")
	if err != nil {
		fmt.Println("Error getting accept-bom flag: ", err)
		return
	}

	if len(args) == 0 && !all {
		fmt.Println("Please provide a file path")
		return
	}

	if all {
		err = spinner.New().
			Title("Converting all files to UTF-8").
			Action(func() { convertAllFilesEncoding(acceptBOM) }).
			Run()
		if err != nil {
			fmt.Println("Error running spinner: ", err)
		}

		return
	}

	path := args[0]

	err = convertEncoding(path, true, acceptBOM)
	if err != nil {
		fmt.Println("Error converting encoding: ", err)
		return
	}
}

func convertAllFilesEncoding(acceptBOM bool) {
	fmt.Println("Converting all files in the current directory to UTF-8")

	files, err := getAllProgramFiles()
	if err != nil {
		fmt.Println("Error getting all files: ", err)
		return
	}

	for _, file := range files {
		err := convertEncoding(file, false, acceptBOM)
		if err != nil {
			fmt.Println("Error converting encoding: ", err)
			return
		}
	}
}

func convertEncoding(path string, verbose bool, acceptBOM bool) error {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return err
	}

	utf8withBom := len(content) >= 3 &&
		content[0] == 0xEF &&
		content[1] == 0xBB &&
		content[2] == 0xBF

	if utf8withBom {

		if acceptBOM {
			return nil
		}

		content = content[3:]
		err := os.WriteFile(path, content, 0o644)
		if err != nil {
			fmt.Println("Convert UTF-8 with BOM failed. Error writing file: ", err)
			return err
		}

		fmt.Println(path + " converted to UTF-8 from UTF-8 with BOM")

		return nil
	}

	if utf8.Valid(content) {
		if verbose {
			fmt.Println(path + " is UTF-8")
		}
		return nil
	}

	big5Encoder := traditionalchinese.Big5.NewDecoder()

	encoderPairs := []encoderPair{{big5Encoder, "Big5"}}

	for _, pair := range encoderPairs {
		exactEncode, err := convertToUTF8(path, content, pair.encoder)
		if err != nil {
			fmt.Println("Error converting to UTF-8: ", err)
			return err
		}

		if exactEncode {
			fmt.Println(path + " converted to UTF-8 from " + pair.encoding)
			return nil
		}
	}

	return nil
}

type encoderPair struct {
	encoder  transform.Transformer
	encoding string
}

func convertToUTF8(path string, content []byte, decoder transform.Transformer) (bool, error) {
	utf8Content, _, err := transform.Bytes(decoder, content)
	if err != nil {
		return false, nil
	}

	err = os.WriteFile(path, utf8Content, 0o644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return true, err
	}

	return true, nil
}

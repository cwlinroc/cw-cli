/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lfCmd)
	rootCmd.AddCommand(crlfCmd)

	lfCmd.Flags().BoolP("all", "a", false, "Change all files in the current directory to LF")
	crlfCmd.Flags().BoolP("all", "a", false, "Change all files in the current directory to CRLF")
}

var lfCmd = &cobra.Command{
	Use:   "lf",
	Short: "change all files change line endings to LF",
	Long:  `change all files change line endings to LF. For example: cw lf`,
	Args:  cobra.RangeArgs(0, 1),
	Run:   lfCmdRun,
}

var crlfCmd = &cobra.Command{
	Use:   "crlf",
	Short: "change all files change line endings to CRLF",
	Long:  `change all files change line endings to CRLF. For example: cw crlf`,
	Args:  cobra.RangeArgs(0, 1),
	Run:   crlfCmdRun,
}

//------------------------------------------------------------

func lfCmdRun(cmd *cobra.Command, args []string) {

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
		spinner.New().
			Title("Converting all files to LF").
			Action(func() {
				files, err := getAllProgramFiles()

				if err != nil {
					fmt.Println(err)
					return
				}

				for _, file := range files {
					err = fileToLf(file)

					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}).
			Run()

		return
	}

	path := args[0]

	err = fileToLf(path)

	if err != nil {
		fmt.Println("Error converting encoding: ", err)
		return
	}
}

func crlfCmdRun(cmd *cobra.Command, args []string) {

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
		spinner.New().
			Title("Converting all files to CRLF").
			Action(func() {
				files, err := getAllProgramFiles()

				if err != nil {
					fmt.Println(err)
					return
				}

				for _, file := range files {
					err = fileToCrlf(file)

					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}).
			Run()

		return
	}

	path := args[0]

	err = fileToCrlf(path)

	if err != nil {
		fmt.Println("Error converting encoding: ", err)
		return
	}
}

func fileToLf(file string) error {
	source, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	source = crlf_to_lf(source)
	err = os.WriteFile(file, source, 0644)

	if err != nil {
		return err
	}

	return nil
}

func fileToCrlf(file string) error {
	source, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	source = lf_to_crlf(source)
	err = os.WriteFile(file, source, 0644)

	if err != nil {
		return err
	}

	return nil
}

//------------------------------------------------------------

func lf_to_crlf(source []byte) []byte {
	lfIndex := get_lf_index(source)

	if len(lfIndex) == 0 {
		return source
	}

	sourceLen := len(source)
	indexLen := len(lfIndex)

	newLen := sourceLen + indexLen

	newSource := make([]byte, newLen)

	for i := 0; i < indexLen; i++ {
		if i == 0 {
			copy(newSource, source[:lfIndex[i]])
		} else {
			copy(newSource[lfIndex[i-1]+i+1:], source[lfIndex[i-1]+1:lfIndex[i]])
		}

		newSource[lfIndex[i]+i] = '\r'
		newSource[lfIndex[i]+i+1] = '\n'
	}

	copy(newSource[lfIndex[indexLen-1]+indexLen+1:], source[lfIndex[indexLen-1]+1:])

	return newSource
}

func crlf_to_lf(source []byte) []byte {
	crlfIndex := get_crlf_index(source)

	if len(crlfIndex) == 0 {
		return source
	}

	sourceLen := len(source)
	indexLen := len(crlfIndex)

	newLen := sourceLen - indexLen

	newSource := make([]byte, newLen)

	for i := 0; i < indexLen; i++ {
		if i == 0 {
			copy(newSource, source[:crlfIndex[i]])
		} else {
			copy(newSource[crlfIndex[i-1]-i+2:], source[crlfIndex[i-1]+2:crlfIndex[i]])
		}

		newSource[crlfIndex[i]-i] = '\n'
	}

	copy(newSource[crlfIndex[indexLen-1]-indexLen+2:], source[crlfIndex[indexLen-1]+2:])

	return newSource
}

func get_lf_index(source []byte) []int {
	sourceLen := len(source)

	if sourceLen == 0 {
		return nil
	}

	lfIndex := make([]int, 0, 10)

	cr := false

	for i := 0; i < sourceLen; i++ {

		len := utf8_len(source[i])

		if len > 1 {
			i += len - 1
			cr = false
			continue
		}

		if source[i] == '\r' {
			cr = true
			continue
		}

		if source[i] == '\n' && !cr {
			lfIndex = append(lfIndex, i)
		}

		cr = false
	}

	return lfIndex

}

func get_crlf_index(source []byte) []int {
	sourceLen := len(source)

	if sourceLen == 0 {
		return nil
	}

	crlfIndex := make([]int, 0, 10)

	for i := 0; i < sourceLen; i++ {

		len := utf8_len(source[i])

		if len > 1 {
			i += len - 1
			continue
		}

		if source[i] == '\r' {
			if i+1 < sourceLen && source[i+1] == '\n' {
				crlfIndex = append(crlfIndex, i)
			}
			i++
		}
	}

	return crlfIndex
}

func utf8_len(b byte) int {
	if b < 0x80 {
		return 1
	} else if b < 0xE0 {
		return 2
	} else if b < 0xF0 {
		return 3
	} else {
		return 4
	}
}

//------------------------------------------------------------

func getAllProgramFiles() ([]string, error) {
	var files []string

	root, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {

			skipedDirects := [...]string{
				".git", ".vscode", "node_modules", "bin", "obj",
				"packages", "dist", "build", "out", "target",
				"__pycache__", ".idea", ".vs", ".nuxt",
			}

			for _, direct := range skipedDirects {
				if direct == info.Name() {
					return filepath.SkipDir
				}
			}

			return nil
		}

		txtFileExtensions := [...]string{
			".py", ".cs", ".go", ".java", ".js", ".ts",
			".cshtml", ".razor", ".csproj", ".sln",
			".html", ".css", ".json", ".xml", ".yaml",
			".yml", ".md", ".txt", ".sh", ".bat", ".cmd",
			".gitignore", ".dockerignore", ".editorconfig",
			".vue", ".php", ".rb", ".cpp", ".h", ".hpp",
			".jsx", ".tsx", ".sql", ".swift", ".svelte",
			".mod", ".sum",
		}

		fileExt := filepath.Ext(path)

		isProgramFile := false

		for _, ext := range txtFileExtensions {
			if ext == fileExt {
				isProgramFile = true
				break
			}
		}

		if !isProgramFile {
			return nil
		}

		files = append(files, path)

		return nil
	})

	return files, err
}

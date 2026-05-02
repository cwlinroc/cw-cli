/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cw/internal/template"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addCsCmd)
	addCmd.AddCommand(addRazorCmd)
	addCmd.AddCommand(addCodeCmd)
	addCmd.AddCommand(addPageCmd)
	addCmd.AddCommand(addEditorConfigCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new file under current directory",
	Long:  `Add a new file under current directory. For example: cw add class`,
}

//------------------------------------------------------------

var addCsCmd = &cobra.Command{
	Use:   "cs",
	Short: "Add a new c# class file under current directory",
	Long:  `Add a new c# class file under current directory. For example: cw add class`,
	Run:   addCsCmdRun,
}

func addCsCmdRun(cmd *cobra.Command, args []string) {

	targetDirect, err := getTargetDirect()

	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	fileName, err := getFileName()

	if err != nil {
		fmt.Println("Error getting file name: ", err)
		return
	}

	err = template.GenerateCs(targetDirect, fileName)

	if err != nil {
		fmt.Println("Error generating class: ", err)
		return
	}
}

var addRazorCmd = &cobra.Command{
	Use:   "razor",
	Short: "Add a new razor page file under current directory",
	Long:  `Add a new razor page file under current directory. For example: cw add razor`,
	Run:   addRazorCmdRun,
}

func addRazorCmdRun(cmd *cobra.Command, args []string) {

	targetDirect, err := getTargetDirect()

	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	fileName, err := getFileName()

	if err != nil {
		fmt.Println("Error getting file name: ", err)
		return
	}

	err = template.GenerateRazor(targetDirect, fileName)

	if err != nil {
		fmt.Println("Error generating razor page: ", err)
		return
	}
}

var addCodeCmd = &cobra.Command{
	Use:   "code",
	Short: "Add a new vscode workspace under current directory",
	Long:  `Add a new vscode workspace under current directory. For example: cw add code`,
	Run:   addCodeCmdRun,
}

func addCodeCmdRun(cmd *cobra.Command, args []string) {

	targetDirect, err := getTargetDirect()

	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	fileName, err := getFileName()

	if err != nil {
		fmt.Println("Error getting file name: ", err)
		return
	}

	err = template.GenerateCode(targetDirect, fileName)

	if err != nil {
		fmt.Println("Error generating code workspace: ", err)
		return
	}
}

var addPageCmd = &cobra.Command{
	Use:   "page",
	Short: "Add a new svelte page file under current directory",
	Long:  `Add a new svelte page file under current directory. For example: cw add page`,
	Run:   addPageCmdRun,
}

func addPageCmdRun(cmd *cobra.Command, args []string) {

	targetDirect, err := getTargetDirect()

	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	err = template.GeneratePage(targetDirect)

	if err != nil {
		fmt.Println("Error generating page: ", err)
		return
	}
}

var addEditorConfigCmd = &cobra.Command{
	Use:   "editorconfig",
	Short: "Add a new .editorconfig file under current directory",
	Long:  `Add a new .editorconfig file under current directory. For example: cw add editorconfig`,
	Run:   addEditorConfigCmdRun,
}

func addEditorConfigCmdRun(cmd *cobra.Command, args []string) {
	targetDirect, err := getTargetDirect()

	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	err = template.GenerateEditorConfig(targetDirect)

	if err != nil {
		fmt.Println("Error generating .editorconfig file: ", err)
		return
	}
}

//------------------------------------------------------------

func getTargetDirect() (targetDirect string, err error) {
	currentDir, err := os.Getwd()

	if err != nil {
		return "", errors.New("Error getting current directory " + err.Error())
	}

	dir := currentDir

	for {
		fmt.Println("Directory: ", dir)

		items, err := os.ReadDir(dir)

		if err != nil {
			return "", errors.New("Error reading directory " + err.Error())
		}

		var entryNames []string
		for _, v := range items {
			if v.IsDir() {
				entryNames = append(entryNames, v.Name())
			}

		}
		entryOptions := make([]huh.Option[string], len(entryNames))

		for i, v := range entryNames {
			entryOptions[i] = huh.NewOption(v, v)
		}

		baseDir := filepath.Base(dir)

		var huhOptions []huh.Option[string]

		if baseDir == "" || baseDir == "/" || baseDir == "\\" {
			huhOptions = append(huh.NewOptions("<current directory>", "<mkdir>"), entryOptions...)
		} else {
			huhOptions = append(huh.NewOptions("<current directory>", "<mkdir>", "../"), entryOptions...)
		}

		var result string

		err = huh.NewSelect[string]().
			Title("Select File Path").
			Options(huhOptions...).
			Value(&result).
			WithTheme(hunTheme).
			Run()

		if err != nil {
			return "", errors.New("Error getting directory name " + err.Error())
		}

		println("Result: ", result)

		if result == "<current directory>" {
			return dir, nil
		} else if result == "../" {
			dir = filepath.Dir(dir)
		} else if result == "<mkdir>" {
			var newDir string

			err := huh.NewInput().
				Title("New Directory Name").
				Value(&newDir).
				WithTheme(hunTheme).
				Run()

			if err != nil {
				return "", errors.New("Error getting new directory name " + err.Error())
			}

			dir = filepath.Join(dir, newDir)

			err = os.Mkdir(dir, 0755)

			if err != nil {
				return "", errors.New("Error creating directory " + err.Error())
			}
		} else if result == "<exit>" {
			return "", errors.New("user exited")
		} else {
			dir = filepath.Join(dir, result)
		}
	}
}

func getFileName() (fileName string, err error) {

	var _fileName string
	// get file name with huh prompt
	{
		for _fileName == "" {

			err := huh.NewInput().
				Title("File Name").
				Value(&_fileName).
				WithTheme(hunTheme).
				Run()

			if err != nil {
				return "", errors.New("Error getting file name " + err.Error())
			}

			_fileName = strings.TrimSpace(_fileName)

			if _fileName == "" {
				fmt.Println("File name cannot be empty")
			}
		}
	}

	return _fileName, nil
}

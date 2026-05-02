/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cw/internal/template"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(touchCmd)

	touchCmd.AddCommand(touchCsCmd)
	touchCmd.AddCommand(touchRazorCmd)
	touchCmd.AddCommand(touchCodeCmd)
	touchCmd.AddCommand(touchPageCmd)
	touchCmd.AddCommand(touchEditorConfigCmd)
}

var touchCmd = &cobra.Command{
	Use:   "touch",
	Short: "Add a new c# class file under current directory",
	Long:  `Add a new c# class file under current directory. For example: cw touch class`,
}

//------------------------------------------------------------

var touchCsCmd = &cobra.Command{
	Use:   "cs",
	Short: "Add a new c# class file under current directory",
	Long:  `Add a new c# class file under current directory. For example: cw touch cs TestClass`,
	Run:   touchCsCmdRun,
	Args:  cobra.RangeArgs(0, 1),
}

func touchCsCmdRun(cmd *cobra.Command, args []string) {

	targetDirect, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	var fileName string

	if len(args) > 0 && args[0] != "" {
		fileName = args[0]
	} else {
		fileName, err = getFileName()

		if err != nil {
			fmt.Println("Error getting file name: ", err)
			return
		}
	}

	if len(fileName) > 3 && fileName[len(fileName)-3:] == ".cs" {
		fileName = fileName[:len(fileName)-3]
	}

	err = template.GenerateCs(targetDirect, fileName)

	if err != nil {
		fmt.Println("Error generating class: ", err)
		return
	}
}

var touchRazorCmd = &cobra.Command{
	Use:   "razor",
	Short: "Add a new razor page file under current directory",
	Long:  `Add a new razor page file under current directory. For example: cw touch razor`,
	Run:   touchRazorCmdRun,
	Args:  cobra.RangeArgs(0, 1),
}

func touchRazorCmdRun(cmd *cobra.Command, args []string) {

	targetDirect, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	var fileName string

	if len(args) > 0 && args[0] != "" {
		fileName = args[0]
	} else {
		fileName, err = getFileName()
		if err != nil {
			fmt.Println("Error getting file name: ", err)
			return
		}
	}

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

var touchCodeCmd = &cobra.Command{
	Use:   "code",
	Short: "Add a new vscode workspace under current directory",
	Long:  `Add a new vscode workspace under current directory. For example: cw touch code`,
	Run:   touchCodeCmdRun,
}

func touchCodeCmdRun(cmd *cobra.Command, args []string) {

	targetDirect, err := os.Getwd()

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

var touchPageCmd = &cobra.Command{
	Use:   "page",
	Short: "Add a new svelte page file under current directory",
	Long:  `Add a new svelte page file under current directory. For example: cw touch page`,
	Run:   touchPageCmdRun,
}

func touchPageCmdRun(cmd *cobra.Command, args []string) {

	targetDirect, err := os.Getwd()

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

var touchEditorConfigCmd = &cobra.Command{
	Use:   "editorconfig",
	Short: "Add a new .editorconfig file under current directory",
	Long:  `Add a new .editorconfig file under current directory. For example: cw touch editorconfig`,
	Run:   touchEditorConfigCmdRun,
}

func touchEditorConfigCmdRun(cmd *cobra.Command, args []string) {

	targetDirect, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	err = template.GenerateEditorConfig(targetDirect)

	if err != nil {
		fmt.Println("Error generating editorconfig: ", err)
		return
	}
}

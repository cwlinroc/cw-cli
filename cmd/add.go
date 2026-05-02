/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"cw/internal/file"
	"cw/internal/template"

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
	targetDirect, err := file.PickDir(huhTheme)
	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	fileName, err := file.GetName(args, huhTheme)
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
	targetDirect, err := file.PickDir(huhTheme)
	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	fileName, err := file.GetName(args, huhTheme)
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
	targetDirect, err := file.PickDir(huhTheme)
	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	fileName, err := file.GetName(args, huhTheme)
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
	targetDirect, err := file.PickDir(huhTheme)
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
	targetDirect, err := file.PickDir(huhTheme)
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

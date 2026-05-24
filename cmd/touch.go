/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"cw/internal/file"
	"cw/internal/template"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(touchCmd)

	touchCmd.AddCommand(touchCsCmd)
	touchCmd.AddCommand(touchRazorCmd)
	touchCmd.AddCommand(touchCodeCmd)
	touchCmd.AddCommand(touchPageCmd)
	touchCmd.AddCommand(touchEditorConfigCmd)
	touchCmd.AddCommand(touchGitignoreCmd)

	touchGitignoreCmd.Flags().BoolVar(&touchGitignoreNet, "net", false, "Include .NET ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignoreDotnet, "dotnet", false, "Include .NET ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignoreGo, "go", false, "Include Go ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignoreGolang, "golang", false, "Include Go ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignoreJava, "java", false, "Include Java ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignoreJs, "js", false, "Include JavaScript/TypeScript ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignoreTs, "ts", false, "Include JavaScript/TypeScript ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignoreNode, "node", false, "Include JavaScript/TypeScript ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignoreNpm, "npm", false, "Include JavaScript/TypeScript ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignorePnpm, "pnpm", false, "Include JavaScript/TypeScript ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignorePy, "py", false, "Include Python ignore rules")
	touchGitignoreCmd.Flags().BoolVar(&touchGitignorePython, "python", false, "Include Python ignore rules")
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

var (
	touchGitignoreNet    bool
	touchGitignoreDotnet bool
	touchGitignoreGo     bool
	touchGitignoreGolang bool
	touchGitignoreJava   bool
	touchGitignoreJs     bool
	touchGitignoreTs     bool
	touchGitignoreNode   bool
	touchGitignoreNpm    bool
	touchGitignorePnpm   bool
	touchGitignorePy     bool
	touchGitignorePython bool
)

var touchGitignoreCmd = &cobra.Command{
	Use:   "gitignore",
	Short: "Add a new .gitignore file under current directory",
	Long:  `Add a new .gitignore file under current directory. Supports --net, --go, --java, --js, --py flags.`,
	Run:   touchGitignoreCmdRun,
}

func touchGitignoreCmdRun(cmd *cobra.Command, args []string) {
	targetDirect, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	opts := template.GitignoreOptions{
		Dotnet: touchGitignoreNet || touchGitignoreDotnet,
		Go:     touchGitignoreGo || touchGitignoreGolang,
		Java:   touchGitignoreJava,
		JsTs:   touchGitignoreJs || touchGitignoreTs || touchGitignoreNode || touchGitignoreNpm || touchGitignorePnpm,
		Python: touchGitignorePy || touchGitignorePython,
	}

	// If no flags are specified, default to including all supported languages.
	if !opts.Dotnet && !opts.Go && !opts.Java && !opts.JsTs && !opts.Python {
		opts.Dotnet = true
		opts.Go = true
		opts.Java = true
		opts.JsTs = true
		opts.Python = true
	}

	err = template.GenerateGitignore(targetDirect, opts)
	if err != nil {
		fmt.Println("Error generating gitignore: ", err)
		return
	}
}

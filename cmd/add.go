/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"cw/internal/file"
	"cw/internal/template"

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
	addCmd.AddCommand(addGitignoreCmd)

	addGitignoreCmd.Flags().BoolVar(&addGitignoreNet, "net", false, "Include .NET ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignoreDotnet, "dotnet", false, "Include .NET ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignoreGo, "go", false, "Include Go ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignoreGolang, "golang", false, "Include Go ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignoreJava, "java", false, "Include Java ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignoreJs, "js", false, "Include JavaScript/TypeScript ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignoreTs, "ts", false, "Include JavaScript/TypeScript ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignoreNode, "node", false, "Include JavaScript/TypeScript ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignoreNpm, "npm", false, "Include JavaScript/TypeScript ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignorePnpm, "pnpm", false, "Include JavaScript/TypeScript ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignorePy, "py", false, "Include Python ignore rules")
	addGitignoreCmd.Flags().BoolVar(&addGitignorePython, "python", false, "Include Python ignore rules")
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

var (
	addGitignoreNet    bool
	addGitignoreDotnet bool
	addGitignoreGo     bool
	addGitignoreGolang bool
	addGitignoreJava   bool
	addGitignoreJs     bool
	addGitignoreTs     bool
	addGitignoreNode   bool
	addGitignoreNpm    bool
	addGitignorePnpm   bool
	addGitignorePy     bool
	addGitignorePython bool
)

var addGitignoreCmd = &cobra.Command{
	Use:   "gitignore",
	Short: "Add a new .gitignore file under current directory",
	Long:  `Add a new .gitignore file under current directory. Supports --net, --go, --java, --js, --py flags.`,
	Run:   addGitignoreCmdRun,
}

func addGitignoreCmdRun(cmd *cobra.Command, args []string) {
	targetDirect, err := file.PickDir(huhTheme)
	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	opts := template.GitignoreOptions{
		Dotnet: addGitignoreNet || addGitignoreDotnet,
		Go:     addGitignoreGo || addGitignoreGolang,
		Java:   addGitignoreJava,
		JsTs:   addGitignoreJs || addGitignoreTs || addGitignoreNode || addGitignoreNpm || addGitignorePnpm,
		Python: addGitignorePy || addGitignorePython,
	}

	// If no flags are specified, prompt with a MultiSelect checkbox UI
	if !opts.Dotnet && !opts.Go && !opts.Java && !opts.JsTs && !opts.Python {
		selected := []string{"dotnet", "go", "java", "jsts", "python"}
		err = huh.NewMultiSelect[string]().
			Title("Select ecosystems to ignore:").
			Options(
				huh.NewOption("Dotnet (.NET)", "dotnet"),
				huh.NewOption("Go", "go"),
				huh.NewOption("Java", "java"),
				huh.NewOption("JS/TS (Node/NPM/PNPM)", "jsts"),
				huh.NewOption("Python", "python"),
			).
			Value(&selected).
			WithTheme(huhTheme).
			Run()
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		for _, s := range selected {
			switch s {
			case "dotnet":
				opts.Dotnet = true
			case "go":
				opts.Go = true
			case "java":
				opts.Java = true
			case "jsts":
				opts.JsTs = true
			case "python":
				opts.Python = true
			}
		}
	}

	err = template.GenerateGitignore(targetDirect, opts)
	if err != nil {
		fmt.Println("Error generating gitignore: ", err)
		return
	}
}

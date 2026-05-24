//go:build !linux

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func init() {
	rootCmd.AddCommand(pathCmd)
}

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Copy the current directory path to the clipboard",
	Long:  `Copy the absolute path of the current working directory to the clipboard.`,
	Args:  cobra.NoArgs,
	Run:   pathCmdRun,
}

func pathCmdRun(cmd *cobra.Command, args []string) {
	err := clipboard.Init()
	if err != nil {
		fmt.Printf("Error accessing clipboard: %v\n", err)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	clipboard.Write(clipboard.FmtText, []byte(dir))
	fmt.Printf("Copied: %s\n", dir)
}

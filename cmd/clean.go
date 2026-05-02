package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Recursively remove build, cache, and dependency directories",
	Long:  `Recursively deletes common build, cache, and dependency directories (such as .cache, bin, obj, node_modules, publish) from the current directory and all subdirectories. Useful for cleaning up project artifacts and restoring a clean workspace.`,
	Args:  cobra.NoArgs,
	Run:   cleanCmdRun,
}

var cleanUpDirectNames = []string{
	".cache",
	"bin",
	"obj",
	"node_modules",
	"publish",
}

func cleanCmdRun(cmd *cobra.Command, args []string) {
	targetDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	fmt.Printf("Cleaning up cache directories in: %s\n", targetDir)

	var removedDirs []string
	var removedCount int64

	err = filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			dirName := filepath.Base(path)
			if slices.Contains(cleanUpDirectNames, dirName) {
				fmt.Printf("Removing: %s\n", path)
				if err := os.RemoveAll(path); err != nil {
					fmt.Printf("Error removing %s: %v\n", path, err)
				} else {
					removedDirs = append(removedDirs, path)
					removedCount++
				}
				return filepath.SkipDir
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory tree: %v\n", err)
		return
	}

	if removedCount > 0 {
		fmt.Printf("\nCleaned up %d directories:\n", removedCount)
		for _, dir := range removedDirs {
			fmt.Printf("  - %s\n", dir)
		}
	} else {
		fmt.Println("No cache directories found to clean up.")
	}
}

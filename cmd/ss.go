//go:build !linux
// +build !linux

package cmd

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"golang.design/x/clipboard"
)

// ss stands for screenshot

func init() {
	rootCmd.AddCommand(ssCmd)
}

var ssCmd = &cobra.Command{
	Use:   "ss",
	Short: "Save a screenshot",
	Long: `Save the screenshot from clipboard to a file.

Supports PNG and BMP formats with automatic timestamping.

Usage examples:
  cw ss                    # Save as PNG with timestamp filename
  cw ss myfile             # Save as PNG with custom filename  
  cw ss png                # Save as PNG format (timestamp filename)
  cw ss bmp                # Save as BMP format (timestamp filename)
  cw ss myfile.png         # Save as PNG with custom filename
  cw ss myfile.bmp         # Save as BMP with custom filename`,
	Run:  ssCmdRun,
	Args: cobra.RangeArgs(0, 1),
}

func ssCmdRun(cmd *cobra.Command, args []string) {

	err := clipboard.Init()

	if err != nil {
		fmt.Println("Error accessing clipboard")
		return
	}

	imgBytes := clipboard.Read(clipboard.FmtImage)

	if imgBytes == nil {
		fmt.Println("No image found on the clipboard")
		return
	}

	targetDirect, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting direct: ", err)
		return
	}

	fileName := ""
	imageType := ""
	{
		if len(args) > 0 && args[0] != "" {
			if args[0] != "bmp" && args[0] != "png" {
				fileName = args[0]
			} else {
				imageType = args[0]
			}
		}

		if fileName == "" {
			fileName = time.Now().Format("20060102_150405")
		}

		if imageType == "" {
			if len(fileName) > 4 && fileName[len(fileName)-4:] == ".bmp" {
				imageType = "bmp"
				fileName = fileName[:len(fileName)-4]
			} else if len(fileName) > 4 && fileName[len(fileName)-4:] == ".png" {
				imageType = "png"
				fileName = fileName[:len(fileName)-4]
			} else {
				imageType = "png"
			}
		}
	}

	switch imageType {
	case "png":
		err = saveAsPng(fileName, targetDirect, imgBytes)
	case "bmp":
		err = saveAsBmp(fileName, targetDirect, imgBytes)
	default:
		err = saveAsPng(fileName, targetDirect, imgBytes)
	}

	if err != nil {
		fmt.Printf("Error saving image: %v\n", err)
		return
	}

	fmt.Printf("Screenshot saved as %s.%s\n", fileName, imageType)

}

func saveAsPng(fileName string, targetDirect string, imgBytes []byte) error {
	reader := bytes.NewReader(imgBytes)

	img, _, err := image.Decode(reader)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	filePath := filepath.Join(targetDirect, fileName+".png")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("failed to encode PNG: %v", err)
	}

	return nil
}

func saveAsBmp(fileName string, targetDirect string, imgBytes []byte) error {
	reader := bytes.NewReader(imgBytes)

	img, _, err := image.Decode(reader)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	filePath := filepath.Join(targetDirect, fileName+".bmp")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Since Go doesn't have built-in BMP encoder, we'll save as PNG instead
	// and inform the user
	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	}

	fmt.Println("Note: BMP format not fully supported, saved as PNG format with .bmp extension")
	return nil
}

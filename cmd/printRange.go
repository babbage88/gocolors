/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	print256Fg    bool
	printAnsi16FG bool
)

// printRangeCmd represents the printRange command
var printRangeCmd = &cobra.Command{
	Use:   "print",
	Short: "Print a range of colors",
	Long:  `Prints the x256 and ANSI16 Foreground colors.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		if printAnsi16FG {
			printAnsi16Fg()
		}
		if print256Fg {
			printXterm256Fg()
		}
	},
}

func init() {
	rootCmd.AddCommand(printRangeCmd)

	printRangeCmd.PersistentFlags().BoolVarP(&printAnsi16FG, "ansi16", "a", true, "Print the ANSI 16 range of foreground colors")
	printRangeCmd.PersistentFlags().BoolVarP(&print256Fg, "xterm256", "x", true, "Print the xterm-256 froeground color")

}

func printXterm256Fg() {
	printSeperatorHorizontal("#")
	fmt.Println("Xterm-256 Color Spectrum (Foreground Colors):")
	printSeperatorHorizontal("#")

	for i := 0; i < 256; i++ {
		// Print color code with foreground color
		fmt.Printf("\033[38;5;%dm %3d \033[0m", i, i)

		// Print 6 colors per row for better readability
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	printSeperatorHorizontal("#")
	fmt.Println()
	fmt.Print("Done.")
	fmt.Println()
}

// printANSI16 prints the standard ANSI 16 foreground colors.
func printAnsi16Fg() {
	printSeperatorHorizontal("#")
	fmt.Println("ANSI 16 Colors (Foreground):")
	printSeperatorHorizontal("#")

	// Standard colors (0-7)
	for i := 0; i < 8; i++ {
		colorId := 30 + i
		fmt.Printf("ANSI Code: %dm", colorId)
		fmt.Printf("\033[%dm %3d Normal  \033[0m\n", colorId, colorId)
		fmt.Printf("ANSI Code: 01;%dm", colorId)
		fmt.Printf("\033[01;%dm %3d Intense \033[0m\n", colorId, colorId)
	}
	fmt.Println()

	// Bright colors (8-15) use 90-97 instead of 30-37
	for i := 0; i < 8; i++ {
		brightId := 90 + i
		fmt.Printf("ANSI Code: %dm", brightId)
		fmt.Printf("\033[%dm %3d Normal  \033[0m\n", brightId, brightId)
		fmt.Printf("ANSI Code: 01;%dm", brightId)
		fmt.Printf("\033[01;%dm %3d Intense  \033[0m\n", brightId, brightId)
	}
	printSeperatorHorizontal("#")
	fmt.Println()
	fmt.Println("Done.")
	fmt.Println()

}

func buildNString(count int, pchar string) string {
	var pChars strings.Builder

	for i := 0; i < count; i++ {
		pChars.WriteString(pchar)
	}

	return pChars.String()
}

func printSeperatorHorizontal(sepChar string) string {
	var sepChars strings.Builder
	width, _, err := term.GetSize(0)
	if err != nil {
		log.Printf("Error determing terminal size.")
		return err.Error()
	}
	for i := 0; i < width; i++ {
		sepChars.WriteString(sepChar)
	}

	fmt.Printf("%s\n", sepChars.String())
	return sepChars.String()

}

/*
func printHeading(heightRatio int, msg string, sepChar string) string {
	var headerChars strings.Builder
	width, height, err := term.GetSize(0)

	headerHeightFloat := math.RoundToEven(float64(height /  heightRatio))

	headerHeightInt := math.RoundToEven(float64(headerHeightFloat))
	headerMidpoint := math.RoundToEven(float64(headerHeight) / 2.0)
	if err != nil {
		log.Printf("Error determing terminal size.")
		return err.Error()
	}

	fullBar := buildNString(width)
	for i := 0; i < headerHeight; i++ {

	}

	return headerChars.String()
}
*/

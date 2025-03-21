/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/babbage88/gocolors/pkg/ansi16"
	"github.com/babbage88/gocolors/pkg/xterm256"
	"github.com/spf13/cobra"
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
			ansi16.PrintAnsi16Fg()
			return
		}
		if print256Fg {
			xterm256.PrintXterm256Fg()
			return
		}

		ansi16.PrintAnsi16Fg()
	},
}

func init() {
	rootCmd.AddCommand(printRangeCmd)

	printRangeCmd.PersistentFlags().BoolVarP(&printAnsi16FG, "ansi16", "a", false, "Print the ANSI 16 range of foreground colors")
	printRangeCmd.PersistentFlags().BoolVarP(&print256Fg, "xterm256", "x", false, "Print the xterm-256 froeground color")

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

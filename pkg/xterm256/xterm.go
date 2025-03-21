package xterm256

import (
	"fmt"

	"github.com/babbage88/gocolors/internal/header"
)

func PrintXterm256Fg() {
	printXterm256Fg()
}

func printXterm256Fg() {
	header.PrintSeperatorHorizontal("#")
	fmt.Println("Xterm-256 Color Spectrum (Foreground Colors):")
	header.PrintSeperatorHorizontal("#")

	for i := 0; i < 256; i++ {
		// Print color code with foreground color
		fmt.Printf("\033[38;5;%dm %3d \033[0m", i, i)

		// Print 6 colors per row for better readability
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	header.PrintSeperatorHorizontal("#")
	fmt.Println()
	fmt.Print("Done.")
	fmt.Println()
}

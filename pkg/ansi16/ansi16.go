package ansi16

import (
	"fmt"

	"github.com/babbage88/gocolors/internal/header"
)

func PrintAnsi16Fg() {
	printAnsi16Fg()

}

// printANSI16 prints the standard ANSI 16 foreground colors.
func printAnsi16Fg() {
	header.PrintSeperatorHorizontal("#")
	fmt.Println("ANSI 16 Colors (Foreground):")
	header.PrintSeperatorHorizontal("#")

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
	header.PrintSeperatorHorizontal("#")
	fmt.Println()
	fmt.Println("Done.")
	fmt.Println()

}

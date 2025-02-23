package main

import (
	"flag"
	"fmt"
)

func printAnsi256Fg() {
	fmt.Println("Xterm-256 Color Spectrum (Foreground Colors):")

	for i := 0; i < 256; i++ {
		// Print color code with foreground color
		fmt.Printf("\033[38;5;%dm %3d \033[0m", i, i)

		// Print 6 colors per row for better readability
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}

	fmt.Println("Done!")
}

// printANSI16 prints the standard ANSI 16 foreground colors.
func printAnsi16Fg() {
	fmt.Println("ANSI 16 Colors (Foreground):")

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
}

func main() {
	// Define command-line flags
	ansi16Flag := flag.Bool("ansi16", true, "Display only ANSI 16 colors")
	xterm256Flag := flag.Bool("xterm256", false, "Display only Xterm-256 colors")
	flag.Parse()

	if *ansi16Flag {
		printAnsi16Fg()
	}
	if *xterm256Flag {
		printAnsi256Fg()
	}
}

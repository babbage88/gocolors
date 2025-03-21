package header

import (
	"fmt"
	"log"
	"math"
	"strings"

	"golang.org/x/term"
)

func PrintSeperatorHorizontal(sepChar string) string {
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

func BuildSeperatorHorizontal(sepChar string) string {
	var sepChars strings.Builder
	width, _, err := term.GetSize(0)
	if err != nil {
		log.Printf("Error determing terminal size.")
		return err.Error()
	}
	for i := 0; i < width; i++ {
		sepChars.WriteString(sepChar)
	}

	sepChar.WriteString("\n")
	return sepChars.String()
}

func BuildNString(count int, pchar string) string {
	var pChars strings.Builder

	for i := 0; i < count; i++ {
		pChars.WriteString(pchar)
	}

	return pChars.String()
}

func PrintHeaderMessage(msg string, sepChar string) string {
	var headerMessage strings.Builder

	width, height, err := term.GetSize(0)
	if err != nil {
		log.Printf("Error determing terminal size.")
		return err.Error()
	}
	heightFloat := math.RoundToEven(float64(height) / 4.0)
	headerHeightInt := int(heightFloat)
	midWidth := int(math.RoundToEven(float64(width) / 2.0))
	msgStart := midWidth - len(msg)
	hdrBodyHeight := headerHeightInt - 2
	hdrBodyHalf := int(math.RoundToEven(float64(hdrBodyHeight) / 2.0))
	headerMessage.WriteString(BuildSeperatorHorizontal(sepChar))

	for i := 0; i < hdrBodyHalf; i++ {
		headerMessage.WriteString(sepChar)
		for w := 0; w < width-2; i++ {
			headerMessage.WriteString(" ")
		}
		headerMessage.WriteString(sepChar)
		headerMessage.WriteString("\n")
	}

	headerMessage.WriteString("#")
	headerMessage.WriteString(BuildNString(msgStart, " "))
	headerMessage.WriteString(msg)

	return headerMessage.String()
}

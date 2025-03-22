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

	sepChars.WriteString("\n")
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
	fmt.Printf("width %d\n", width)

	fmt.Printf("midWidth %d\n", midWidth)
	fmt.Printf("len(msg) %d\n", len(msg))

	msgStart := midWidth - len(msg)
	fmt.Printf("messageStart %d\n", msgStart)

	hdrBodyHeight := headerHeightInt - 2
	hdrBodyHalf := int(math.RoundToEven(float64(hdrBodyHeight) / 2.0))
	msgLinePaddingRight := midWidth - 2

	headerMessage.WriteString(BuildSeperatorHorizontal(sepChar))

	buildHeaderBody(hdrBodyHalf, width, sepChar, &headerMessage)

	headerMessage.WriteString("#")
	headerMessage.WriteString(BuildNString(msgStart, " "))
	headerMessage.WriteString(msg)
	headerMessage.WriteString(BuildNString(msgLinePaddingRight, " "))
	headerMessage.WriteString(sepChar)
	headerMessage.WriteString("\n")

	buildHeaderBody(hdrBodyHalf, width, sepChar, &headerMessage)
	headerMessage.WriteString(BuildSeperatorHorizontal(sepChar))

	return headerMessage.String()
}

func buildHeaderBody(height int, width int, sepChar string, headerBuilder *strings.Builder) {
	for x := 0; x < height; x++ {
		headerBuilder.WriteString(sepChar)
		for y := 0; y < width-2; y++ {
			headerBuilder.WriteString(" ")
		}
		headerBuilder.WriteString(sepChar)
		headerBuilder.WriteString("\n")
	}
}

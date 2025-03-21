package header

import (
	"fmt"
	"log"
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

func BuildNString(count int, pchar string) string {
	var pChars strings.Builder

	for i := 0; i < count; i++ {
		pChars.WriteString(pchar)
	}

	return pChars.String()
}

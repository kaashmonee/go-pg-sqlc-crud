package util

import (
	"bytes"
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func JsonStringPrettyPrint(input []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, input, "", "  "); err != nil {
		log.Printf("Failed to pretty print JSON: %v", err)
		return "", err
	}

	// Create color objects
	keyColor := color.New(color.FgGreen).SprintFunc()
	stringColor := color.New(color.FgYellow).SprintFunc()
	numberColor := color.New(color.FgCyan).SprintFunc()
	boolColor := color.New(color.FgRed).SprintFunc()

	// Get the indented string
	colored := prettyJSON.String()

	// Colorize the output using regex
	colored = regexp.MustCompile(`"(\\.|[^"])*"`).ReplaceAllStringFunc(colored, func(s string) string {
		if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
			// Check if it's a key (followed by :)
			if strings.Contains(s+`"`, `":`) {
				return keyColor(s)
			}
			return stringColor(s)
		}
		return s
	})
	colored = regexp.MustCompile(`:\s*(-?\d+\.?\d*)`).ReplaceAllString(colored, ": "+numberColor(`$1`))
	colored = regexp.MustCompile(`:\s*(true|false|null)`).ReplaceAllString(colored, ": "+boolColor(`$1`))

	return colored + "\n", nil
}

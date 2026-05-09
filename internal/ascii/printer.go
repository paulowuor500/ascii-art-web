package ascii

import (
	"errors"
	"strings"
)

func GenerateASCII(input, banner string) (string, error) {
	filePath := "banners/" + banner + ".txt"

	font, err := loadBanner(filePath)
	if err != nil {
		return "", err
	}

	var result strings.Builder

	// Split multiline input
	lines := strings.Split(input, "\n")

	for _, line := range lines {

		// ASCII characters are 8 rows tall
		for row := 0; row < 8; row++ {

			// Process each character in the line
			for _, char := range line {

				// Validate printable ASCII range
				if char < 32 || char > 126 {
					return "", errors.New("unsupported character")
				}

				index := int(char - 32)

				// Append the correct ASCII row
				result.WriteString(font[index*9+row])
			}

			// Move to next ASCII row
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

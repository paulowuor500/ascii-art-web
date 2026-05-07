package ascii

import (
	"errors"
	"strings"
)

func GenerateASCII(input, banner string) (string, error) {
	filepath := "banners/" + banner + ".txt"

	font, err := loadBanner(filepath)
	if err != nil {
		return "", err
	}

	var result strings.Builder

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		for i := 0; i < 8; i++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					return "", errors.New("unsupported character")
				}

				index := int(char - 32)

				result.WriteString(font[index*9+i])
			}

			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

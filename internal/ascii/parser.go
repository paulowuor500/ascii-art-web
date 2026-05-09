package ascii

import (
	"os"
	"strings"
)

func loadBanner(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Remove Windows carriage returns
	content := strings.ReplaceAll(string(data), "\r", "")

	lines := strings.Split(content, "\n")

	return lines, nil
}

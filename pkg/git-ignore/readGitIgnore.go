package git_ignore

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadGitIgnore(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines or comments (#)
		if line == "" || gitignoreComments.MatchString(line) {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return lines, nil
}

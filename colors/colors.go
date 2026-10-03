package colors

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadColorsToml(path string) map[string]string {
	result := make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Trying to read:", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			parts := strings.Split(line, "=")

			if len(parts) != 2 {
				log.Fatal("Couldn't split", line, "by = into 2 parts...")
			}

			key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

			result[key] = value
		}
	}

	return result
}

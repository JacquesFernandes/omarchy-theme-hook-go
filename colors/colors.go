package colors

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// TODO: Figure out mapping between om3 and 4

type Colors struct {
	mode,

	accent,
	selection,
	muted, // oma3 color0 normal_black

	background, // oma3 background primary_background
	darkBackground,
	darkerBackground,
	lighterBackground,

	foreground,
	darkForeground,
	lightForeground,
	brightForeground,

	red,
	yellow,
	orange,
	green,
	cyan,
	blue,
	magenta,
	brown,

	brightRed,
	brightYellow,
	brightGreen,
	brightCyan,
	brightBlue,
	brightMagenta string
}

func loadColorsToml(path string) map[string]string {
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

func ParseColorsToml(colorsTomlPath string) Colors {
	file, err := os.Open(colorsTomlPath)
	if err != nil {
		log.Fatal("Trying to read:", colorsTomlPath, err)
	}
	defer file.Close()

	buffer := bufio.NewScanner(file)
	colors := Colors{}
	for buffer.Scan() {
		line := buffer.Text()
		parts := strings.Split(line, "=")

		if len(parts) != 2 {
			log.Fatal("Couldn't split", line, "by = into 2 parts...")
		}

		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

		switch key {
		case "mode":
			colors.mode = value
		case "accent":
			colors.accent = value
		case "selection":
			colors.selection = value
		case "muted":
			colors.muted = value
		case "background":
			colors.background = value
		case "dark_background":
			colors.darkBackground = value
		case "darker_background":
			colors.darkerBackground = value
		case "lighter_background":
			colors.lighterBackground = value
		case "foreground":
			colors.foreground = value
		case "dark_foreground":
			colors.darkForeground = value
		case "light_foreground":
			colors.lightForeground = value
		case "bright_foreground":
			colors.brightForeground = value
		case "red":
			colors.red = value
		case "yellow":
			colors.yellow = value
		case "orange":
			colors.orange = value
		case "green":
			colors.green = value
		case "cyan":
			colors.cyan = value
		case "blue":
			colors.blue = value
		case "magenta":
			colors.magenta = value
		case "brown":
			colors.brown = value
		case "bright_red":
			colors.brightRed = value
		case "bright_yellow":
			colors.brightYellow = value
		case "bright_green":
			colors.brightGreen = value
		case "bright_cyan":
			colors.brightCyan = value
		case "bright_blue":
			colors.brightBlue = value
		case "bright_magenta":
			colors.brightMagenta = value
		}
	}

	return colors
}

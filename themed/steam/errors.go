package steam

import (
	"fmt"
	"strings"
)

type missingDependencyError struct {
	commandName  string
	instructions []string
}

func (e *missingDependencyError) Error() string {
	firstLine := fmt.Sprintf("error: Missing command/executible/script %q", e.commandName)

	if len(e.instructions) == 0 {
		return firstLine
	}

	var builder strings.Builder

	builder.WriteString("\n") // Adds new line to previous
	builder.WriteString("Try doing the following:\n")
	for _, instruction := range e.instructions {
		builder.WriteString(" - " + instruction + "\n")
	}

	builderResult := builder.String()

	return firstLine + builderResult
}

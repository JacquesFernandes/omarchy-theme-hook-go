package steam

import (
	"fmt"
	"strings"
	"testing"
)

func TestMissingDependencyError_onlyCommandGiven(t *testing.T) {
	err := missingDependencyError{
		commandName: "missing-command",
	}

	output := err.Error()

	// Fail if commandName is absent
	if !strings.Contains(output, err.commandName) {
		t.Error("Command name is missing from error:", err)
	}

	// Fail if there is more than one line
	if strings.Contains(output, "\n") {
		t.Error("Error contains a newline when it shouldn't:", err)
	}
}

func TestMissingDependencyError_commandAndInstructionsGiven(t *testing.T) {
	err := missingDependencyError{
		commandName: "missing-command-with-instructions",
		instructions: []string{
			"step one",
			"step two",
		},
	}

	output := err.Error()

	// Fail if commandName is absent
	if !strings.Contains(output, err.commandName) {
		t.Error("Command name is missing from error:", output)
	}

	// Fail if there aren't multiple lines
	if !strings.Contains(output, "\n") {
		t.Error("Error is missing newlines:", output)
	}

	// Fail if instructions are missing
	for _, instruction := range err.instructions {
		t.Run(fmt.Sprintf("check for instruction: %q", instruction), func(rt *testing.T) {
			if !strings.Contains(output, instruction) {
				rt.Error("missing instruction:", instruction)
			}
		})
	}
}

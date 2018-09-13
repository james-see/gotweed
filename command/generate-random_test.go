package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGenerateRandomCommand_implement(t *testing.T) {
	var _ cli.Command = &GenerateRandomCommand{}
}

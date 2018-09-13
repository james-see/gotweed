package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGenerateMarkovCommand_implement(t *testing.T) {
	var _ cli.Command = &GenerateMarkovCommand{}
}

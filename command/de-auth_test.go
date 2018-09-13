package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDeAuthCommand_implement(t *testing.T) {
	var _ cli.Command = &DeAuthCommand{}
}

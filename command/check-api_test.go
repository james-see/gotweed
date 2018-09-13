package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestCheckApiCommand_implement(t *testing.T) {
	var _ cli.Command = &CheckApiCommand{}
}

package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestArchiveUserCommand_implement(t *testing.T) {
	var _ cli.Command = &ArchiveUserCommand{}
}

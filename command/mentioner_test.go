package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestMentionerCommand_implement(t *testing.T) {
	var _ cli.Command = &MentionerCommand{}
}

package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAddToMentionerListCommand_implement(t *testing.T) {
	var _ cli.Command = &AddToMentionerListCommand{}
}

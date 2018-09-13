package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestSearchUserCommand_implement(t *testing.T) {
	var _ cli.Command = &SearchUserCommand{}
}

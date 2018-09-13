package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAssociateUserWithProjectCommand_implement(t *testing.T) {
	var _ cli.Command = &AssociateUserWithProjectCommand{}
}

package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDeleteProjectCommand_implement(t *testing.T) {
	var _ cli.Command = &DeleteProjectCommand{}
}

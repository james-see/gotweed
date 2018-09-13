package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestLoadProjectCommand_implement(t *testing.T) {
	var _ cli.Command = &LoadProjectCommand{}
}

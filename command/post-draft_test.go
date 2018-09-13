package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestPostDraftCommand_implement(t *testing.T) {
	var _ cli.Command = &PostDraftCommand{}
}

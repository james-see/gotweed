package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDeleteDraftCommand_implement(t *testing.T) {
	var _ cli.Command = &DeleteDraftCommand{}
}

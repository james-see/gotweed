package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestListDraftsCommand_implement(t *testing.T) {
	var _ cli.Command = &ListDraftsCommand{}
}

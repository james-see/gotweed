package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAuthorizeCommand_implement(t *testing.T) {
	var _ cli.Command = &AuthorizeCommand{}
}

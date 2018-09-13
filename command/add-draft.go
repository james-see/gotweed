package command

import (
	"strings"
)

type AddDraftCommand struct {
	Meta
}

func (c *AddDraftCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *AddDraftCommand) Synopsis() string {
	return ""
}

func (c *AddDraftCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

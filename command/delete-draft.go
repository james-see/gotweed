package command

import (
	"strings"
)

type DeleteDraftCommand struct {
	Meta
}

func (c *DeleteDraftCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DeleteDraftCommand) Synopsis() string {
	return ""
}

func (c *DeleteDraftCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

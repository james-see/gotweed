package command

import (
	"strings"
)

type PostDraftCommand struct {
	Meta
}

func (c *PostDraftCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *PostDraftCommand) Synopsis() string {
	return ""
}

func (c *PostDraftCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

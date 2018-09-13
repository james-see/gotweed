package command

import (
	"strings"
)

type ListDraftsCommand struct {
	Meta
}

func (c *ListDraftsCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ListDraftsCommand) Synopsis() string {
	return ""
}

func (c *ListDraftsCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

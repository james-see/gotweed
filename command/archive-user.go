package command

import (
	"strings"
)

type ArchiveUserCommand struct {
	Meta
}

func (c *ArchiveUserCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ArchiveUserCommand) Synopsis() string {
	return ""
}

func (c *ArchiveUserCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

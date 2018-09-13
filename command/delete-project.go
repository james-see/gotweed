package command

import (
	"strings"
)

type DeleteProjectCommand struct {
	Meta
}

func (c *DeleteProjectCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DeleteProjectCommand) Synopsis() string {
	return ""
}

func (c *DeleteProjectCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

package command

import (
	"strings"
)

type SetProjectCommand struct {
	Meta
}

func (c *SetProjectCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SetProjectCommand) Synopsis() string {
	return ""
}

func (c *SetProjectCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

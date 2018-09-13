package command

import (
	"strings"
)

type DeAuthCommand struct {
	Meta
}

func (c *DeAuthCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DeAuthCommand) Synopsis() string {
	return ""
}

func (c *DeAuthCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

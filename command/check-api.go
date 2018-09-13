package command

import (
	"strings"
)

type CheckApiCommand struct {
	Meta
}

func (c *CheckApiCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *CheckApiCommand) Synopsis() string {
	return ""
}

func (c *CheckApiCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

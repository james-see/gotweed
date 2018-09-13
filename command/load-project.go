package command

import (
	"strings"
)

type LoadProjectCommand struct {
	Meta
}

func (c *LoadProjectCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *LoadProjectCommand) Synopsis() string {
	return ""
}

func (c *LoadProjectCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

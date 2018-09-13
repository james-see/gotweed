package command

import (
	"strings"
)

type AuthorizeCommand struct {
	Meta
}

func (c *AuthorizeCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *AuthorizeCommand) Synopsis() string {
	return ""
}

func (c *AuthorizeCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

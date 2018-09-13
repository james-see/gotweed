package command

import (
	"strings"
)

type SearchUserCommand struct {
	Meta
}

func (c *SearchUserCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *SearchUserCommand) Synopsis() string {
	return ""
}

func (c *SearchUserCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

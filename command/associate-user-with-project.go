package command

import (
	"strings"
)

type AssociateUserWithProjectCommand struct {
	Meta
}

func (c *AssociateUserWithProjectCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *AssociateUserWithProjectCommand) Synopsis() string {
	return ""
}

func (c *AssociateUserWithProjectCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

package command

import (
	"strings"
)

type AddToMentionerListCommand struct {
	Meta
}

func (c *AddToMentionerListCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *AddToMentionerListCommand) Synopsis() string {
	return ""
}

func (c *AddToMentionerListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

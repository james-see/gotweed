package command

import (
	"strings"
)

type MentionerCommand struct {
	Meta
}

func (c *MentionerCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *MentionerCommand) Synopsis() string {
	return ""
}

func (c *MentionerCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

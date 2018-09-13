package command

import (
	"strings"
)

type GetRandomCommand struct {
	Meta
}

func (c *GetRandomCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GetRandomCommand) Synopsis() string {
	return ""
}

func (c *GetRandomCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

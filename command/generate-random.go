package command

import (
	"strings"
)

type GenerateRandomCommand struct {
	Meta
}

func (c *GenerateRandomCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GenerateRandomCommand) Synopsis() string {
	return ""
}

func (c *GenerateRandomCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

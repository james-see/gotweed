package command

import (
	"strings"
)

type GenerateMarkovCommand struct {
	Meta
}

func (c *GenerateMarkovCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GenerateMarkovCommand) Synopsis() string {
	return ""
}

func (c *GenerateMarkovCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

package bot

import (
	"github.com/micro/micro/bot/command"
)

type sortedCommands struct {
	commands []command.Command
}

func (s sortedCommands) Len() int {
	return len(s.commands)
}

func (s sortedCommands) Less(i, j int) bool {
	return s.commands[i].Name() < s.commands[j].Name()
}

func (s sortedCommands) Swap(i, j int) {
	s.commands[i], s.commands[j] = s.commands[j], s.commands[i]
}

package command

type Command struct {
	name        string
	description string
	handler     func()
}

func New(name string) *Command {
	command := &Command{
		name:    name,
		handler: func() {},
	}

	return command
}

func (command *Command) Name() string {
	return command.name
}

func (command *Command) Description() string {
	return command.description
}

func (command *Command) SetDescription(value string) *Command {
	command.description = value

	return command
}

func (command *Command) Handle(value func()) *Command {
	command.handler = value

	return command
}

func (command *Command) Execute() {
	command.handler()
}

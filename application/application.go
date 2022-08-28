package application

import (
	"fmt"
	"github.com/midonis/action/command"
	"os"
	"path/filepath"
)

type Application struct {
	commands    map[string]*command.Command
	name        string
	description string
	version     string
}

func New() *Application {
	app := &Application{
		version:  "1.0.0",
		commands: make(map[string]*command.Command),
	}

	app.bootstrap()

	return app
}

func (app *Application) Name() string {
	return filepath.Base(os.Args[0])
}

func (app *Application) Description() string {
	return app.description
}

func (app *Application) SetDescription(value string) *Application {
	app.description = value

	return app
}

func (app *Application) Version() string {
	return app.version
}

func (app *Application) SetVersion(value string) *Application {
	app.version = value

	return app
}

func (app *Application) Commands(commands ...*command.Command) *Application {
	for _, command := range commands {
		app.commands[command.Name()] = command
	}

	return app
}

func (app *Application) Run() {
	if app.getCommandName() == "" {
		app.showAvailableCommands()
		return
	}

	cmd := app.commands[app.getCommandName()]

	if cmd == nil {
		println("Could not find command: " + app.getCommandName())
		println()
		app.showAvailableCommands()
		return
	}

	cmd.Execute()
}

func (app *Application) bootstrap() {
	app.name = app.Name()

	listCommand := command.New("list")
	listCommand.SetDescription("Displays all available commands")
	listCommand.Handle(func() {
		app.showAvailableCommands()
	})

	app.Commands(listCommand)
}

func (app *Application) showAvailableCommands() {
	fmt.Printf("  USAGE: %s <command>", app.name)
	println()
	println()
	for _, cmd := range app.commands {
		println("  " + cmd.Name())
		if cmd.Description() != "" {
			println("  - " + cmd.Description())
		}
	}
}

func (app *Application) getCommandName() string {
	if len(os.Args) <= 1 {
		return ""
	}

	return os.Args[1]
}

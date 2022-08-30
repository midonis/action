<p align="center"><img width="100" src="https://github.com/midonis/action/raw/HEAD/art/logo.svg" alt="Logo Midonis Action"></p>

## About Action

Action takes the pain out of building CLI applications from scratch, focusing on your next big idea. It allows you to create beautiful, powerful and stable CLI applications with a lot of flexibility.

## Installation

The implementation of Action is pretty easy. Navigate to your Go project and run the following in your console:
```sh
go get github.com/midonis/action
```

## Example

Here is an example application with two different commands. A `say:hello` and a `say:bye` command:

```go
package main

import (
	"github.com/midonis/action/application"
	"github.com/midonis/action/command"
)

func main() {
	app := application.New()

	app.Commands(
		command.New("say:hello").Handle(func() {
			println("Hello!")
		}),
		command.New("say:bye").Handle(func() {
			println("Bye!")
		}),
		// ...
	)

	app.Run()
}
```

The logic of the specific command goes into the `Handle` function.

## License

Midonis Action is open-sourced software licensed under the [MIT license](LICENSE.md).
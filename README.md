# Bubbletea Test Application

A first look and play with the [Bubbletea terminal user interface](https://github.com/charmbracelet/bubbletea). The following is not meant as a complete guide for beginners and is not affiliated with the official Bubbletea project in any way, it is just my own exploration of the project. If it is useful to you, great! Some of the work done here follows the [official tutorials](https://github.com/charmbracelet/bubbletea/tree/master/tutorials).

## The Bubbletea Project

### The Model Interface
Bubbletea applications consist of a model struct that defines the application state, and three functions associated with that model:
- `Init()`: A function to define the initial state of the application, setting up any state, resources, etc. Init can return a `tea.Cmd` (see below), which is run by Bubbletea once the program has started up.
- `Update()`: A function to handle incoming events and update the model. Update can also return a `tea.Cmd`, for applications that need to have some kind of event loop.
- `View()`: A function to render the model state. View only returns a string, literally the string to be rendered to the terminal.

### `tea.Cmd` and `tea.Msg`

A `tea.Cmd` is just a type definition: 

```go
type Cmd func() Msg
```

i.e. a command is just any function returning exactly one message. A message is again just a simple type definition:

```go
type Msg struct{}
```

So a message is just any struct you'd like. Remember in the Update method we tend to use a type switch for the incoming message. Something like:

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
    switch msg.(type) {
    case ...:
    case ...:
    case ...:
    }
}
```

So we can define new structs such as `MyResponseMsg` and put this in a new case in our type switch, allowing us to cleanly access the fields of our custom defined message without hassle.

Commands are run asynchronously in a goroutine by Bubbletea and, once completed, have the returned message passed to the Update function. This makes commands useful for defining I/O (getting the results of a file read, database read, HTTP request etc...) without having to implement all of the logic for asynchronous responses ourselves. It is a good idea to have all the I/O in commands. We can store the results of I/O in our application state so the View method can be aware of in progress commands.

If you need to parse arguments to a command, instead create a wrapper function that takes those arguments and returns a command.



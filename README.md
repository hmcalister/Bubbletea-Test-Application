# Bubbletea Test Application

A first look and play with the [Bubbletea terminal user interface](https://github.com/charmbracelet/bubbletea). The following is not meant as a complete guide for beginners and is not affiliated with the official Bubbletea project in any way, it is just my own exploration of the project. If it is useful to you, great! Some of the work done here follows the [official tutorials](https://github.com/charmbracelet/bubbletea/tree/master/tutorials).

## The Bubbletea Project

### The Model Interface
Bubbletea applications consist of a model struct that defines the application state, and three functions associated with that model:
- `Init()`: A function to define the initial state of the application, setting up any state, resources, etc. Init can return a `tea.Cmd` (see below), which is run by Bubbletea once the program has started up.
- `Update()`: A function to handle incoming events and update the model. Update can also return a `tea.Cmd`, for applications that need to have some kind of event loop.
- `View()`: A function to render the model state. View only returns a string, literally the string to be rendered to the terminal.


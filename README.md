# SimpleTodo

SimpleTodo is a Todo list manager written in Go. It allows you to add tasks, view tasks, and save your Todo list to a file.

---

## Usage

### Compile
After cloning the repository and navigating into the directory, build the application by running:
```bash
go build -o SimpleTodo ./cmd/SimpleTodo/main.go
```
### Add a Task
Run the app and follow the prompt to enter a new task.

### View All Tasks
Use the `--list` flag to view the current list of tasks:
```bash
./SimpleTodo --list
```

### File Management
- The app saves your tasks to a file named `save.stodo`.
- If the file doesn't exist, it will be created automatically.

# golang examples

## [todo](./todo) - CLI ToDo app in Golang

![](./todo_screen.png)

```
Usage of todo:
  -add
        add a new todo
  -complete int
        mark a todo as completed
  -del int
        delete a todo
  -list
        list all todos
```

using with pipe

```
$ echo "New task" | ./todo -add
```
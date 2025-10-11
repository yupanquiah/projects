package main

import "github.com/yupanquiah/projects/task-tracker/internal/tasks"

func main() {
	tasks.AddTask("Comprar pan")
	tasks.AddTask("Limpiar la casa")
	tasks.MarkTask(1, tasks.StatusInProgress)
	tasks.UpdateTask(2, "Limpiar la casa y sacar la basura")
	tasks.ListTasks("")
}

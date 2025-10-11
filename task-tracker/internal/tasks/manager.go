package tasks

import (
	"errors"
	"fmt"
	"time"
)

func AddTask(description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	if err := saveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", id)
	return nil
}

func ListTasks(filter string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No hay tareas registradas.")
		return nil
	}

	for _, t := range tasks {
		if filter == "" || t.Status == filter {
			fmt.Printf("[%d] %-30s | %-12s | Creado: %s\n",
				t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04"))
		}
	}
	return nil
}

func UpdateTask(id int, newDesc string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = newDesc
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return errors.New("tarea no encontrada")
	}

	if err := saveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task %d updated successfully\n", id)
	return nil
}

func DeleteTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("tarea no encontrada")
	}

	if err := saveTasks(newTasks); err != nil {
		return err
	}

	fmt.Printf("Task %d deleted successfully\n", id)
	return nil
}

func MarkTask(id int, status string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	validStatus := map[string]bool{
		StatusTodo:       true,
		StatusInProgress: true,
		StatusDone:       true,
	}

	if !validStatus[status] {
		return errors.New("estado inválido")
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return errors.New("tarea no encontrada")
	}

	if err := saveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task %d marked as %s\n", id, status)
	return nil
}

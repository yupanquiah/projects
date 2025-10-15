package tasks

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/yupanquiah/projects/task-tracker/internal/ui"
	UITable "github.com/yupanquiah/projects/task-tracker/internal/ui/table"
	"github.com/yupanquiah/projects/task-tracker/internal/utils"
)

var logger = utils.NewLogger()

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		logger.Warn("No tasks found in your list")
		return nil
	}

	var filteredTasks []Task
	switch status {
	case "all":
		filteredTasks = tasks
	case STATUS_TODO:
		for _, task := range tasks {
			if task.Status == STATUS_TODO {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == STATUS_IN_PROGRESS {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case STATUS_DONE:
		for _, task := range tasks {
			if task.Status == STATUS_DONE {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	if len(filteredTasks) == 0 {
		logger.Warn(fmt.Sprintf("No tasks found with status '%s'", status))
		return nil
	}

	taskCount := len(filteredTasks)
	title := fmt.Sprintf("\nTasks (%s) - %d total\n", status, taskCount)
	fmt.Println(ui.Title.Render(title))

	columns := []table.Column{
		{Title: "ID", Width: 5},
		{Title: "Description", Width: 40},
		{Title: "Status", Width: 15},
		{Title: "Updated At", Width: 20},
	}

	rows := []table.Row{}
	for _, task := range filteredTasks {
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", task.ID),
			task.Description,
			string(task.Status),
			task.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	UITable.Create(columns, rows)
	return nil
}

func TaskStatusFromString(status string) TaskStatus {
	switch status {
	case "todo":
		return STATUS_TODO
	case "in-progress":
		return STATUS_IN_PROGRESS
	case "done":
		return STATUS_DONE
	default:
		return "all"
	}
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var newTaskId int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskId = lastTask.ID + 1
	} else {
		newTaskId = 1
	}

	task := NewTask(newTaskId, description)
	tasks = append(tasks, *task)

	formattedId := ui.Id.Render(fmt.Sprintf("(ID: %d)", task.ID))
	logger.Info(fmt.Sprintf("Task added successfully %s", formattedId))

	return WriteTasksToFile(tasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	if len(updatedTasks) == len(tasks) {
		logger.Error(fmt.Sprintf("Task not found (ID: %d)", id))
		return nil
	}

	formattedId := ui.Id.Render(fmt.Sprintf("(ID: %d)", id))
	logger.Info(fmt.Errorf("Task deleted successfully %s", formattedId))
	return WriteTasksToFile(updatedTasks)
}

func UpdateTaskStatus(id int64, status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			switch status {
			case STATUS_TODO:
				task.Status = STATUS_TODO
			case STATUS_IN_PROGRESS:
				task.Status = STATUS_IN_PROGRESS
			case STATUS_DONE:
				task.Status = STATUS_DONE
			}
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		logger.Error(fmt.Errorf("Task not found (ID: %d)", id))
		return nil
	}

	formattedId := ui.Id.Render(fmt.Sprintf("(ID: %d)", id))
	logger.Info(fmt.Sprintf("Task updated successfully: %s", formattedId))
	return WriteTasksToFile(updatedTasks)
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			task.Description = description
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		logger.Error(fmt.Errorf("Task not found (ID: %d)", id))
		return nil
	}

	formattedId := ui.Id.Render(fmt.Sprintf("(ID: %d)", id))
	logger.Info(fmt.Errorf("Task updated successfully: %s", formattedId))
	return WriteTasksToFile(updatedTasks)
}

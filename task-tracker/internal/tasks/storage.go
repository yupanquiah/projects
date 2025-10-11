package tasks

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const filePath = "data/tasks.json"

// ensureDataFileExists: verifica que exista la carpeta y archivo JSON.
func ensureDataFileExists() error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return os.WriteFile(filePath, []byte("[]"), 0644)
	}

	return nil
}

func loadTasks() ([]Task, error) {
	if err := ensureDataFileExists(); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if len(data) == 0 {
		return []Task{}, nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, errors.New("error al leer el archivo de tareas")
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

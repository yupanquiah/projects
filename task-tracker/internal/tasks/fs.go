package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func tasksFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}
	return path.Join(cwd, "data/tasks.json")
}

func ReadTasksFromFile() ([]Task, error) {
	filePath := tasksFilePath()
	dir := path.Dir(filePath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return nil, err
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating file...")
		if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
			fmt.Println("Error creating file:", err)
			return nil, err
		}
		return []Task{}, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	tasks := []Task{}
	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Error decoding file:", err)
		fmt.Println("Resetting file to empty array []...")
		_ = os.WriteFile(filePath, []byte("[]"), 0644)
		return []Task{}, nil
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	filePath := tasksFilePath()
	dir := path.Dir(filePath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(tasks); err != nil {
		fmt.Println("Error encoding file:", err)
		return err
	}

	return nil
}

// 0775 and 0664 permissions are for Unix-based systems. On Windows, these permissions are ignored.

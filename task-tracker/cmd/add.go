package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/tasks"
	"github.com/yupanquiah/projects/task-tracker/internal/utils"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to the task list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(cmd, args)
		},
	}
	return cmd
}

func RunAddTaskCmd(cmd *cobra.Command, args []string) error {
	logger := utils.NewLogger()

	if len(args) == 0 {
		err := errors.New("task description is required")
		logger.Error(err.Error())
		_ = cmd.Usage()

		return nil
	}

	description := args[0]
	return tasks.AddTask(description)
}

package cmd

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/tasks"
	"github.com/yupanquiah/projects/task-tracker/internal/utils"
)

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task",
		Long: `Update a task by providing the task ID and the new status
    Example:
    task-tracker update 1 'new description'
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskCmd(cmd, args)
		},
	}

	return cmd
}

func RunUpdateTaskCmd(cmd *cobra.Command, args []string) error {
	logger := utils.NewLogger()

	if len(args) != 2 {
		err := errors.New("please provide a task id and the new description")
		logger.Error(err.Error())
		_ = cmd.Usage()

		return nil
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		return err
	}

	newDescription := args[1]
	return tasks.UpdateTaskDescription(taskIDInt, newDescription)
}

func NewStatusDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark a task as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(cmd, args, tasks.STATUS_DONE)
		},
	}
	return cmd
}

func NewStatusInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark a task as in-progress",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(cmd, args, tasks.STATUS_IN_PROGRESS)
		},
	}
	return cmd
}

func NewStatusTodoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-todo",
		Short: "Mark a task as todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(cmd, args, tasks.STATUS_TODO)
		},
	}
	return cmd
}

func RunUpdateStatusCmd(cmd *cobra.Command, args []string, status tasks.TaskStatus) error {
	if len(args) == 0 {

		logger := utils.NewLogger()
		err := errors.New("task ID is required")
		logger.Error(err.Error())
		_ = cmd.Usage()

		return nil
	}

	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	return tasks.UpdateTaskStatus(id, status)
}

package cmd

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/tasks"
	"github.com/yupanquiah/projects/task-tracker/internal/utils"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Long: `Delete a task by providing the task ID

    Example:
    task delete 1
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(cmd, args)
		},
	}

	return cmd
}

func RunDeleteTaskCmd(cmd *cobra.Command, args []string) error {
	logger := utils.NewLogger()

	if len(args) != 1 {
		err := errors.New("please provide a task ID")
		logger.Error(err.Error())
		_ = cmd.Usage()

		return nil
	}

	taskID := args[0]
	taskIDInt, err := strconv.ParseInt(taskID, 10, 32)
	if err != nil {
		logger.Error("invalid task ID\n")
		_ = cmd.Usage()

		return nil
	}

	return tasks.DeleteTask(taskIDInt)
}

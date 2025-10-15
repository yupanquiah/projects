package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/tasks"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long: `List all tasks. You can filter tasks by status

    Example:
    task-tracker list todo
    task-tracker list in-progress
    task-tracker list done
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTaskCmd(args)
		},
	}
	return cmd
}

func RunListTaskCmd(args []string) error {
	if len(args) > 0 {
		status := tasks.TaskStatus(args[0])
		return tasks.ListTasks(status)
	}

	return tasks.ListTasks("all")
}

package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/tasks"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Elimina una tarea por su ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		return tasks.DeleteTask(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

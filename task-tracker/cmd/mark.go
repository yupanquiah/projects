package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/tasks"
)

var markCmd = &cobra.Command{
	Use:   "mark",
	Short: "Cambia el estado de una tarea",
}

var markDoneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Marca una tarea como completada",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		return tasks.MarkTask(id, tasks.StatusDone)
	},
}

var markProgressCmd = &cobra.Command{
	Use:   "in-progress [id]",
	Short: "Marca una tarea como en progreso",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		return tasks.MarkTask(id, tasks.StatusInProgress)
	},
}

func init() {
	markCmd.AddCommand(markDoneCmd)
	markCmd.AddCommand(markProgressCmd)
	rootCmd.AddCommand(markCmd)
}

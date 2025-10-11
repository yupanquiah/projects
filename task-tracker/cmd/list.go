package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/tasks"
)

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "Lista las tareas (todas o filtradas por estado)",
	Long: `Muestra todas las tareas o solo aquellas con un estado específico.
Estados válidos: todo, in-progress, done.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filter := ""
		if len(args) > 0 {
			filter = args[0]
		}
		return tasks.ListTasks(filter)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

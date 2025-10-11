package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/tasks"
)

var addCmd = &cobra.Command{
	Use:   "add [descripcion]",
	Short: "Agrega una nueva tarea",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		description := strings.Join(args, " ")
		if err := tasks.AddTask(description); err != nil {
			return err
		}
		fmt.Println("✅ Tarea agregada correctamente.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

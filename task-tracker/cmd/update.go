package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/tasks"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [nueva descripción]",
	Short: "Actualiza la descripción de una tarea",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("el ID debe ser un número entero")
		}
		newDesc := strings.Join(args[1:], " ")
		return tasks.UpdateTask(id, newDesc)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

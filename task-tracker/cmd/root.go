package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yupanquiah/projects/task-tracker/internal/ui"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   ui.Cmd.Render("task"),
		Short: ui.Cmd.Render("Task Tracker") + " CLI gestiona tus tareas desde la terminal",
	}

	cmd.AddCommand(AddCmd())
	cmd.AddCommand(ListCmd())
	cmd.AddCommand(DeleteCmd())
	cmd.AddCommand(UpdateCmd())
	cmd.AddCommand(StatusDoneCmd())
	cmd.AddCommand(StatusInProgressCmd())
	cmd.AddCommand(StatusTodoCmd())

	return cmd
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Task Tracker CLI - gestiona tus tareas desde la terminal",
	Long: `Task Tracker es una herramienta de línea de comandos
para gestionar tus tareas (todo, in-progress, done)
utilizando un archivo JSON local.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usa 'task-cli --help' para ver los comandos disponibles.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

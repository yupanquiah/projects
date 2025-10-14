package main

import (
	"os"

	"github.com/yupanquiah/projects/task-tracker/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(0)
	}
}

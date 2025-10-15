package utils

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/yupanquiah/projects/task-tracker/internal/ui"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stderr)

	logger.SetReportTimestamp(false)
	logger.SetReportCaller(false)
	logger.SetFormatter(log.TextFormatter)

	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = ui.Error
	styles.Levels[log.InfoLevel] = ui.Success
	styles.Levels[log.WarnLevel] = ui.Warning

	logger.SetStyles(styles)

	return logger
}

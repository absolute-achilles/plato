package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

// TODO: Maybe use lumberjack logger for log rotation later

type Config struct {
	Level     slog.Level // slog.LevelDebug, Info, Warn, Error
	LogDir    string     // e.g. "logs"
	Filename  string     // e.g. "app.log"
	AddSource bool       // include file:line in log output
}

// New creates a logger that writes to both stdout and a log file.
// Returns the logger and a cleanup function to close the file.
func New(cfg Config) (logger *slog.Logger, cleanup func(), err error) {
	// create the log directory if it doesn't exist
	if err := os.MkdirAll(cfg.LogDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("logger: failed to create log dir: %w", err)
	}

	logPath := filepath.Join(cfg.LogDir, cfg.Filename)

	// open or create the log file — append mode
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("logger: failed to open log file: %w", err)
	}

	// write to both stdout AND the file simultaneously
	multiWriter := io.MultiWriter(os.Stdout, file)

	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		Level:     cfg.Level,
		AddSource: cfg.AddSource,
	})

	logger = slog.New(handler)

	cleanup = func() {
		file.Close()
	}

	return logger, cleanup, nil
}

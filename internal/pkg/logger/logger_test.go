package logger

import (
	"github.com/paul-ss/pgram-backend/internal/pkg/config"
	"testing"
)

func TestNew(t *testing.T) {
	config.InitTestConfig(&config.Config{
		Logger: config.Logger{
			Level:    "",
			Filepath: "test.log",
			JSON:     false,
			Stdout:   false,
		},
	})

	//assert.Panics(t, func() { _ = newLogger() })

}

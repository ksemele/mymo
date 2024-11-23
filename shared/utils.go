package shared

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// Logger — exportable logger that can be used in other modules
var Logger zerolog.Logger

// initializes zerolog with the specified settings
func InitializeLogger(debug bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

// WIP custom wrapper experiments
func LogDebug(message string, fields map[string]interface{}) {
	event := log.Debug()
	for key, value := range fields {
		event = event.Interface(key, value)
	}
	event.Msg(message)
}

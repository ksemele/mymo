package main

import (
	"flag"
	s "github.com/ksemele/mymo/shared"
)

func main() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	s.InitializeLogger(*debug)

	s.Logger.Debug().Msg("This message appears only when log level set to Debug")
	s.Logger.Info().Msg("This message appears when log level set to Debug or Info")

	if e := s.Logger.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}

	// custom wrapper usage
	s.LogDebug("User login attempt", map[string]interface{}{
		"username": "johndoe",
		"success":  true,
	})

	s.LogDebug("System check complete", nil)
}

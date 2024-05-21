package log

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

var once sync.Once

func initLogger() {
	logger = zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339},
	).Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
}

func GetLogger() *zerolog.Logger {
	once.Do(func() {
		initLogger()
	})

	return &logger
}

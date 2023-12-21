package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func Logs() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logFileName := "logs.json"
	file, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
	} else {
		log.Output(zerolog.ConsoleWriter{Out: file})
	}

	logger := zerolog.New(file).With().Timestamp().Logger()

	fmt.Println("Testing")
	logger.Info().Msg("Testing")
}

package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logs() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	fmt.Println("Testing")
	log.Info().Msg("Testing")
}

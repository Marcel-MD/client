package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Marcel-MD/client/domain"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config()
	domain.SetConfig(cfg)

	for i := 0; i < cfg.NrOfClients; i++ {
		client := domain.NewClient()
		go client.Run()
	}
}

func config() domain.Config {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	file, err := os.Open("config/cfg.json")
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening menu.json")
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var cfg domain.Config
	json.Unmarshal(byteValue, &cfg)

	return cfg
}

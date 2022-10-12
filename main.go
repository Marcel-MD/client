package main

import (
	"encoding/json"
	"io"
	"os"
	"sync/atomic"
	"time"

	"github.com/Marcel-MD/client/domain"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config()
	domain.SetConfig(cfg)

	for {
		if int(atomic.LoadInt64(&domain.NrOfClients)) < cfg.NrOfClients {
			client := domain.NewClient()
			atomic.AddInt64(&domain.NrOfClients, 1)
			go client.Run()
		} else {
			time.Sleep(time.Duration(100*cfg.TimeUnit) * time.Millisecond)
		}
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

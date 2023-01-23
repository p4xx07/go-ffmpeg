package configuration

import "github.com/caarlos0/env/v6"

type Configuration struct {
	FfmpegPath  string `env:"FFMPEG_PATH"`
	FfprobePath string `env:"FFPROBE_PATH"`
	IsDebug     bool   `env:"IS_DEBUG"`
}

func GetConfiguration() Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		panic("failed to read configuration")
	}
	return configuration
}

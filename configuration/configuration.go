package configuration

import "github.com/caarlos0/env/v6"

type Configuration struct {
	FfmpegPath  string `env:"FFMPEG_PATH"`
	FfprobePath string `env:"FFPROBE_PATH"`
	IsDebug     bool   `env:"IS_DEBUG"`
	TestInput   string `env:"TEST_INPUT"`
	TestOutput  string `env:"TEST_OUTPUT"`
}

func GetConfiguration() *Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		panic("failed to read configuration")
	}
	return &configuration
}

package configuration

type Configuration struct {
	FfmpegPath  string
	FfprobePath string
	IsDebug     bool
}

func (c *Configuration) Init() {
	c.FfmpegPath = "/opt/homebrew/bin/ffmpeg"
	c.FfprobePath = "/opt/homebrew/bin/ffprobe"
	c.IsDebug = true
}

package configuration

type Configuration struct {
	FfmpegPath  string
	FfprobePath string
}

func (c *Configuration) Init() {
	c.FfmpegPath = "/opt/homebrew/bin/ffmpeg"
	c.FfprobePath = "/opt/homebrew/bin/ffprobe"
}

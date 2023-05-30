package ffprobe

import (
	"encoding/json"
	"github.com/Paxx-RnD/go-ffmpeg/configuration"
	"os/exec"
)

type IFfprobe interface {
	GetAudioStreams(input string) (*Probe, error)
	GetVideoStreams(input string) (*Probe, error)
	GetProbe(args ...string) (*Probe, error)
}

func NewFfProbe(configuration *configuration.Configuration, headers []string) IFfprobe {
	return &ffprobe{
		configuration: configuration,
		headers:       headers,
	}
}

type ffprobe struct {
	configuration *configuration.Configuration
	headers       []string
}

func (p *ffprobe) GetAudioStreams(input string) (*Probe, error) {
	args := []string{
		"-select_streams", "a",
		input,
	}
	return p.GetProbe(args...)
}

func (p *ffprobe) GetVideoStreams(input string) (*Probe, error) {
	args := []string{
		"-select_streams", "v",
		input,
	}
	return p.GetProbe(args...)
}

func (p *ffprobe) GetProbe(args ...string) (*Probe, error) {
	cmd := exec.Command(p.configuration.FfprobePath)
	cArgs := []string{
		"-show_streams",
		"-show_format",
		"-print_format", "json",
	}

	cmd.Args = append(cmd.Args, cArgs...)
	cmd.Args = append(cmd.Args, args...)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var probe Probe
	err = json.Unmarshal(output, &probe)
	if err != nil {
		return nil, err
	}

	return &probe, nil
}

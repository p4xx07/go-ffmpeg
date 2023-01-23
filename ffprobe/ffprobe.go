package ffprobe

import (
	"encoding/json"
	"github.com/Paxx-RnD/go-ffmpeg/configuration"
	"os/exec"
)

type Ffprobe struct {
	Configuration configuration.Configuration
	Headers       []string
}

func (p *Ffprobe) GetAudioStreams(input string) (*Probe, error) {
	args := []string{
		"-select_streams", "a",
		input,
	}
	return p.GetStreams(args...)
}

func (p *Ffprobe) GetVideoStreams(input string) (*Probe, error) {
	args := []string{
		"-select_streams", "v",
		input,
	}
	return p.GetStreams(args...)
}

func (p *Ffprobe) GetStreams(args ...string) (*Probe, error) {
	cArgs := []string{
		"-print_format", "json",
		"-show_streams",
	}

	cArgs = append(cArgs, args...)
	return p.GetProbe(cArgs...)
}

func (p *Ffprobe) GetProbe(args ...string) (*Probe, error) {
	cmd := exec.Command(p.Configuration.FfprobePath)
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

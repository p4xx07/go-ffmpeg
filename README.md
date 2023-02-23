# Go Ffmpeg

[![Go build](https://github.com/Paxx-RnD/go-ffmpeg/actions/workflows/go-build.yml/badge.svg)](https://github.com/Paxx-RnD/go-ffmpeg/actions/workflows/go-build.yml)
[![Go Test workflow](https://github.com/Paxx-RnD/go-ffmpeg/actions/workflows/go-test.yml/badge.svg)](https://github.com/Paxx-RnD/go-ffmpeg/actions/workflows/go-test.yml)

## Install
```
go get github.com/Paxx-RnD/go-ffmpeg
```

### Simple Example 
```go
func main(){
   headers := []string{
		"-y",
		"-hide_banner",
		"-loglevel", "info",
	}
    
    configuration := configuration.NewConfiguration(
		"usr/bin/ffmpeg",
		"usr/bin/ffprobe",
		false)
        
	f := ffmpeg.NewFfmpeg(configuration, headers)
    
    args := f.
        Input("/path/to/video.mp4").
        Output("/path/to/output.mp4").
        Build()

    f.Run(args)
}
```
### Example with bitrate and codecs
```go
func main(){
    headers := []string{
		"-y",
		"-hide_banner",
		"-loglevel", "info",
	}
    
    configuration := configuration.NewConfiguration(
		"usr/bin/ffmpeg",
		"usr/bin/ffprobe",
		false)
        
	f := ffmpeg.NewFfmpeg(configuration, headers)

    args := f.
        Input("/path/to/video.mp4").
        BitrateVideo(common_bitrates.VideoBitrate100K).
        BitrateAudio(common_bitrates.AudioBitrate128K).
        CodecVideo(codec_video.LIBX264).
        CodecAudio(codec_audio.AAC).
        Output("/path/to/output.mp4").
        Build()

    f.Run(args)
}
```
### Example with Filter Complex
```go
func main(){
    headers := []string{
		"-y",
		"-hide_banner",
		"-loglevel", "info",
	}
    
    configuration := configuration.NewConfiguration(
		"usr/bin/ffmpeg",
		"usr/bin/ffprobe",
		false)
        
	f := ffmpeg.NewFfmpeg(configuration, headers)

    args := f.
        Input("/path/to/video.mp4").
        CodecVideo(codec_video.LIBX264).
        CodecAudio(codec_audio.AAC).
        FilterGraph().
        Fps("0:v", 15, "fps1").
        Scale("fps1", 100, 100, "scale").
        Map("scale").
        Output("/path/to/output.mp4").
        Build()
        
    f.Run(args)
}
```

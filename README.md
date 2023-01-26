# Go Ffmpeg

[![Go build](https://github.com/Paxx-RnD/go-ffmpeg/actions/workflows/go-build.yml/badge.svg)](https://github.com/Paxx-RnD/go-ffmpeg/actions/workflows/go-build.yml)
[![Go Test workflow](https://github.com/Paxx-RnD/go-ffmpeg/actions/workflows/go-test.yml/badge.svg)](https://github.com/Paxx-RnD/go-ffmpeg/actions/workflows/go-test.yml)

## Install
```
go get github.com/Paxx-RnD/go-ffmpeg
```

### Example
```go
func main(){
    f := ffmpeg.Ffmpeg{
            Configuration: configuration.Configuration{
                FfmpegPath: "/usr/bin/ffmpeg"
            }
        }

    args := f.
        Input("/path/to/video.mp4").
        CodecVideo(codec_video.LIBX264).
        CodecAudio(codec_audio.AAC).
        FilterGraph().
        Fps("0:v", 15, "fps1").
        Scale("fps1", 100, 100, "scale").
        Map("scale").
        Output("path/to/output.mp4").
        Build()
        
    f.Run(args)
}
```

---
layout: default
title: FFProbe
nav_order: 2
---

# FFProbe 

The FFprobe module in go-ffmpeg provides a way to retrieve information about media files using the FFprobe command. It allows you to extract information such as codecs, streams, format, and metadata from a media file.

## Usage
Here is an example of how to use the FFprobe module:

```go
package main

import (
	ffmpegConfiguration "github.com/Paxx-RnD/go-ffmpeg/configuration"
	"github.com/Paxx-RnD/go-ffmpeg/ffprobe" 
)


func main(){
   //
	ffConfig := ffmpegConfiguration.GetConfiguration()
	ffprobeService := ffprobe.Ffprobe{Configuration: ffConfig}
   //
}

```

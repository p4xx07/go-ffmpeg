---
layout: default
title: Home
nav_order: 1
description: "GO FFmpegDocs"
permalink: /
last_modified_date: 2023-01-23T17:00:00+0000
---

# GO FFmpeg Official Documentation

go-ffmpeg is a Go library for running FFmpeg commands. It provides a simple and convenient way to execute FFmpeg commands and retrieve their output.

## Installation
Use `go get` to install the library:

```bash
go get github.com/Paxx-RnD/go-ffmpeg
```

## Usage
Here is an example of how to use the library:

```go

package main

import (
	"fmt"

	ffmpeg "github.com/Paxx-RnD/go-ffmpeg"
)

func main() {
	// Create a new FFmpeg instance
	cmd := ffmpeg.New("/usr/local/bin/ffmpeg")

	// Set the input file
	input := ffmpeg.Input("input.mp4")

	// Set the output file
	output := ffmpeg.Output("output.mp3")

	// Set the audio codec
	audio := ffmpeg.AudioCodec("libmp3lame")

	// Build the command
	cmd.Input(input).Output(output).AudioCodec(audio)

	// Run the command
	output, err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output
	fmt.Println(output)
}
```

## Contributing
We welcome contributions to this project. If you want to contribute, please fork the repository and submit a pull request with your changes.

## License
go-ffmpeg is licensed under the MIT License. See LICENSE for more information.

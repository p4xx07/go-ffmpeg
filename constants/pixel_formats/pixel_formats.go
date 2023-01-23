package pixel_formats

type PixelFormat string

const (
	YUV420P  PixelFormat = "yuv420p"
	YUV422P  PixelFormat = "yuv422p"
	YUV444P  PixelFormat = "yuv444p"
	RGB      PixelFormat = "rgb"
	RGBA     PixelFormat = "rgba"
	GRAY     PixelFormat = "gray"
	PAL8     PixelFormat = "pal8"
	YUVA420P PixelFormat = "yuva420p"
	YUVA422P PixelFormat = "yuva422p"
	YUVA444P PixelFormat = "yuva444p"
)

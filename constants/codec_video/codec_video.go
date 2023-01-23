package codec_video

type CodecVideo string

const (
	LIBKVAZAAR  CodecVideo = "libkvazaar"
	LIBOPENH264 CodecVideo = "libopenh264"
	LIBTHEORA   CodecVideo = "libtheora"
	LIBXVID     CodecVideo = "libxvid"
	LIBVPX      CodecVideo = "libvpx"
	LIBVPX_VP9  CodecVideo = "libvpx-vp9"
	LIBWEBP     CodecVideo = "libwebp"
	LIBX264     CodecVideo = "libx264"
	LIBX265     CodecVideo = "libx265"
	H264_NVENC  CodecVideo = "h264_nvenc"
	H264_CUVID  CodecVideo = "h264_cuvid"
	HEVC_NVENC  CodecVideo = "hevc_nvenc"
	H264_AMF    CodecVideo = "h264_amf"
	HVEC_AMF    CodecVideo = "hvec_amf"
)

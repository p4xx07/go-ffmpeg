package codec_audio

type CodecAudio string

const (
	AAC        CodecAudio = "aac"
	AC3        CodecAudio = "ac3"
	FLAC       CodecAudio = "flac"
	LIBFDK_AAC CodecAudio = "libfdk_aac"
	LIBMP3LAME CodecAudio = "libmp3lame"
	LIBOPUS    CodecAudio = "libopus"
	LIBSHINE   CodecAudio = "libshine"
	LIBTWOLAME CodecAudio = "libtwolame"
	LIBVORBIS  CodecAudio = "libvorbis"
	LIBWAVPACK CodecAudio = "libwavpack"
)

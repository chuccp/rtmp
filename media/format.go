package media

type VideoPacket struct {
}

type VideoInfo struct {
	AvCodeId AVCodeId
	Height   uint32
	Width    uint32
	Duration float64
	Framerate uint32
}

func NewVideoInfo(AvCodeId AVCodeId) *VideoInfo {
	return &VideoInfo{AvCodeId:AvCodeId}
}
package media

type VideoPacket struct {
}

type VideoInfo struct {
	AvCodeId AVCodeId
	Height   uint32
	Width    uint32
}

func NewVideoInfo(AvCodeId AVCodeId) *VideoInfo {
	return &VideoInfo{AvCodeId:AvCodeId}
}
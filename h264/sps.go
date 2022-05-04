package h264

import (
	"github.com/chuccp/rtmp/util"
)

type ProfileIdc byte

const (
	BaseLine ProfileIdc = 66
	Main                = 77
	Extended            = 88
	High                = 100
	High10              = 10
	High422             = 122
	High444             = 144
)

type SPS struct {
	profileIdc        ProfileIdc
	constraintFlag    byte
	levelIdc          byte
	data              []byte
	seqParameterSetId uint64
	chromaFormatIdc   uint64

	separateColourPlaneFlag bool

	bitDepthLumaMinus8              uint64
	bitDepthChromaMinus8            uint64
	qpprimeYZeroTransformBypassFlag bool

	seqScalingMatrixPresentFlag bool
	seqScalingListPresentFlag   []bool
	log2MaxFrameNumMinus4       uint64
	picOrderCntType             uint64
	log2MaxPicOrderCntLsbMinus4 uint64

	deltaPicOrderAlwaysZeroFlag    bool
	offsetForNonRefPic             int64
	offsetForTopToBottomField      int64
	numRefFramesInPicOrderCntCycle uint64
	offsetForRefFrame              []int64

	numRefFrames                   uint64
	gapsInFrameNumValueAllowedFlag bool
	picWidthInMbsMintus1           uint64
	picHeightInMapUnitsMinus1      uint64

	frameMbsOnlyFlag bool

	mbAdaptiveFrameFieldFlag bool
	direct88InferenceFlag    bool
	frameCroppingFlag        bool
	frameCropLeftOffset      uint64
	frameCropRightOffset     uint64
	frameCropTopOffset       uint64
	frameCropBottomOffset    uint64

	vuiParametersPresentFlag bool



}

func ParseSPS(naul *NAUL) (*SPS, error) {
	return (&SPS{data: naul.data}).init()
}

func (s *SPS) init() (*SPS, error) {
	s.profileIdc = ProfileIdc(s.data[0])
	s.constraintFlag = s.data[1]
	s.levelIdc = s.data[2]
	gd := util.NewGolombDecode(s.data[3:])
	seqParameterSetId, err := gd.ReadUGolomb()
	if err != nil {
		return nil, err
	}
	s.seqParameterSetId = seqParameterSetId
	if s.profileIdc == 100 || s.profileIdc == 110 || s.profileIdc == 122 ||
		s.profileIdc == 244 || s.profileIdc == 44 || s.profileIdc == 83 ||
		s.profileIdc == 86 || s.profileIdc == 118 || s.profileIdc == 128 {
		chromaFormatIdc, err := gd.ReadUGolomb()
		if err != nil {
			return nil, err
		} else {
			s.chromaFormatIdc = chromaFormatIdc
		}
		if chromaFormatIdc == 3 {
			s.separateColourPlaneFlag, err = gd.ReadBit()
			if err != nil {
				return nil, err
			}
		}
		s.bitDepthLumaMinus8, err = gd.ReadUGolomb()
		if err != nil {
			return nil, err
		}
		s.bitDepthChromaMinus8, err = gd.ReadUGolomb()
		if err != nil {
			return nil, err
		}
		s.qpprimeYZeroTransformBypassFlag, err = gd.ReadBit()
		if err != nil {
			return nil, err
		}
		s.seqScalingMatrixPresentFlag, err = gd.ReadBit()
		if err != nil {
			return nil, err
		}
		if s.seqScalingMatrixPresentFlag {
			max := 12
			if s.chromaFormatIdc != 3 {
				max = 8
			}
			s.seqScalingListPresentFlag = make([]bool, max)
			for i := 0; i < max; i++ {
				s.seqScalingListPresentFlag[i], err = gd.ReadBit()
				if err != nil {
					return nil, err
				}
			}
		}
	}
	s.log2MaxFrameNumMinus4, err = gd.ReadUGolomb()
	if err != nil {
		return nil, err
	}
	s.picOrderCntType, err = gd.ReadUGolomb()
	if err != nil {
		return nil, err
	}
	if s.picOrderCntType == 0 {
		s.log2MaxPicOrderCntLsbMinus4, err = gd.ReadUGolomb()
		if err != nil {
			return nil, err
		}
	} else if s.picOrderCntType == 1 {
		s.deltaPicOrderAlwaysZeroFlag, err = gd.ReadBit()
		if err != nil {
			return nil, err
		}
		s.offsetForNonRefPic, err = gd.ReadGolomb()
		if err != nil {
			return nil, err
		}
		s.offsetForTopToBottomField, err = gd.ReadGolomb()
		if err != nil {
			return nil, err
		}
		s.numRefFramesInPicOrderCntCycle, err = gd.ReadUGolomb()
		if err != nil {
			return nil, err
		}
		num := int(s.numRefFramesInPicOrderCntCycle)
		s.offsetForRefFrame = make([]int64, num)
		for i := 0; i < int(s.numRefFramesInPicOrderCntCycle); i++ {
			s.offsetForRefFrame[i], err = gd.ReadGolomb()
			if err != nil {
				return nil, err
			}
		}

	}
	s.numRefFrames, err = gd.ReadUGolomb()
	if err != nil {
		return nil, err
	}
	s.gapsInFrameNumValueAllowedFlag, err = gd.ReadBit()
	if err != nil {
		return nil, err
	}
	s.picWidthInMbsMintus1, err = gd.ReadUGolomb()
	if err != nil {
		return nil, err
	}
	s.picHeightInMapUnitsMinus1, err = gd.ReadUGolomb()
	if err != nil {
		return nil, err
	}
	s.frameMbsOnlyFlag,err = gd.ReadBit()
	if err!=nil{
		return nil, err
	}
	if !s.frameMbsOnlyFlag{
		s.mbAdaptiveFrameFieldFlag,err = gd.ReadBit()
		if err!=nil{
			return nil, err
		}
	}

	s.direct88InferenceFlag,err  = gd.ReadBit()
	if err!=nil{
		return nil, err
	}
	s.frameCroppingFlag,err =  gd.ReadBit()
	if err!=nil{
		return nil, err
	}
	if s.frameCroppingFlag{
		s.frameCropLeftOffset,err = gd.ReadUGolomb()
		if err!=nil{
			return nil, err
		}
		s.frameCropRightOffset,err = gd.ReadUGolomb()
		if err!=nil{
			return nil, err
		}
		s.frameCropTopOffset,err = gd.ReadUGolomb()
		if err!=nil{
			return nil, err
		}
		s.frameCropBottomOffset,err = gd.ReadUGolomb()
		if err!=nil{
			return nil, err
		}
	}

	s.vuiParametersPresentFlag,err = gd.ReadBit()
	if err!=nil{
		return nil, err
	}

	return s, nil
}
func (s *SPS) ProfileIdc() ProfileIdc {
	return s.profileIdc
}
func (s *SPS) LevelIdc() byte {
	return s.levelIdc
}
func (s *SPS) SeqParameterSetId() uint64 {
	return s.seqParameterSetId
}
func (s *SPS) ChromaFormatIdc() uint64 {
	return s.chromaFormatIdc
}
func (s *SPS) BitDepthLumaMinus8() uint64 {
	return s.bitDepthLumaMinus8
}

func (s *SPS) BitDepthChromaMinus8() uint64 {
	return s.bitDepthChromaMinus8
}

func (s *SPS) SeqScalingMatrixPresentFlag() bool {
	return s.seqScalingMatrixPresentFlag
}
func (s *SPS) PicOrderCntType() uint64 {
	return s.picOrderCntType
}
func (s *SPS) Log2MaxFrameNumMinus4() uint64 {
	return s.log2MaxFrameNumMinus4
}
func (s *SPS) Log2MaxPicOrderCntLsbMinus4() uint64 {
	return s.log2MaxPicOrderCntLsbMinus4
}
func (s *SPS) PicHeightInMapUnitsMinus1() uint64 {
	return s.picHeightInMapUnitsMinus1
}
func (s *SPS) PicWidthInMbsMintus1() uint64 {
	return s.picWidthInMbsMintus1
}
func (s *SPS) NumRefFrames() uint64 {
	return s.numRefFrames
}
func (s *SPS) GapsInFrameNumValueAllowedFlag() bool {
	return s.gapsInFrameNumValueAllowedFlag
}
func (s *SPS) FrameMbsOnlyFlag()bool  {
	return s.frameMbsOnlyFlag
}
func (s *SPS) VuiParametersPresentFlag()bool  {
	return s.vuiParametersPresentFlag
}
func (s *SPS) Direct88InferenceFlag()bool  {
	return s.direct88InferenceFlag
}
func (s *SPS) FrameCroppingFlag()bool  {
	return s.frameCroppingFlag
}
//frameCroppingFlag
//direct88InferenceFlag
//log2MaxPicOrderCntLsbMinus4

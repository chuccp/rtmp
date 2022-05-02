package h264

import "testing"

func TestOOO(t *testing.T) {
	d,err:=Open("C:\\Users\\cooge\\Videos\\123321.h264")
	var num = 0
	if err==nil{
		for{
			num++
			naul,err:=d.ReadNAUL()
			if err==nil{
				println(naul.ToHexString())
				println(naul.NAULType())
				ty:=naul.NAULType()
				if ty==SPSType {

					sps,err:=ParseSPS(naul)
					if err==nil{
						println(sps.ProfileIdc())
						println("sqid",sps.SeqParameterSetId())
						println("chromaFormatIdc",sps.ChromaFormatIdc())
						println("bitDepthLumaMinus8",sps.BitDepthLumaMinus8())
						println("BitDepthChromaMinus8",sps.BitDepthChromaMinus8())
						println("SeqScalingMatrixPresentFlag",sps.SeqScalingMatrixPresentFlag())
						println("Log2MaxFrameNumMinus4",sps.Log2MaxFrameNumMinus4())
						println("PicOrderCntType",sps.PicOrderCntType())
						println("Log2MaxPicOrderCntLsbMinus4",sps.Log2MaxPicOrderCntLsbMinus4())
						println("NumRefFrames",sps.NumRefFrames())
						println("GapsInFrameNumValueAllowedFlag",sps.GapsInFrameNumValueAllowedFlag())
						println("PicWidthInMbsMintus1",sps.PicWidthInMbsMintus1())
						println("PicHeightInMapUnitsMinus1",sps.PicHeightInMapUnitsMinus1())
						println("frameMbsOnlyFlag",sps.FrameMbsOnlyFlag())
						println("direct88InferenceFlag",sps.Direct88InferenceFlag())
						println("frameCroppingFlag",sps.FrameCroppingFlag())
						println("vuiParametersPresentFlag",sps.VuiParametersPresentFlag())

					}

				}
			}else{
				break
			}
			if num>=3{
				break
			}
		}
	}

}

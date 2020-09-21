package ocr

type PointT struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PosT struct{
	LeftTop     PointT `json:"left_top"`
	RigthTop    PointT `json:"right_top"`
	RightBottom PointT `json:"right_bottom"`
	LeftBottom  PointT `json:"left_bottom"`
}

type ImgSizeT struct {
	Width  int `json:"w"`
	Heigth int `json:"h"`
}

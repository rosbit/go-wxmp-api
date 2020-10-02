package img

type (
	Point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	Pos struct{
		LeftTop     Point `json:"left_top"`
		RigthTop    Point `json:"right_top"`
		RightBottom Point `json:"right_bottom"`
		LeftBottom  Point `json:"left_bottom"`
	}

	Crop struct{
		Left   int `json:"crop_left"`
		Top    int `json:"crop_top"`
		Right  int `json:"crop_right"`
		Bottom int `json:"crop_bottom"`
	}

	ImgSize struct {
		Width  int `json:"w"`
		Heigth int `json:"h"`
	}
)

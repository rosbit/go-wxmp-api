// 本接口提供基于小程序的条码/二维码识别的API
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html

package img

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"bytes"
	"fmt"
	"io"
	u "net/url"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type QRCodeRes struct {
	CodeResults []struct{
		TypeName string `json:"type_name"`
		Data string `json:"data"`
		Pos struct{
			LeftTop     Point `json:"left_top"`
			RigthTop    Point `json:"right_top"`
			RightBottom Point `json:"right_bottom"`
			LeftBottom  Point `json:"left_bottom"`
		} `json:"pos"`
	} `json:"code_results"`
	ImgSize struct {
		Width  int `json:"w"`
		Heigth int `json:"h"`
	} `json:"img_size"`
}

func ScanQrcodeImgUrl(cfgName, imgUrl string) (*QRCodeRes, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		v := u.Values{}
		v.Set("img_url", imgUrl)
		v.Set("access_token", accessToken)
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/img/qrcode?%s", v.Encode())
		return
	}

	return scanQrcode(cfgName, genParams)
}

func ScanQrcodeImg(cfgName, fileName string, fp io.Reader) (*QRCodeRes, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = fmt.Sprintf("https://api.weixin.qq.com/cv/img/qrcode?access_token=%s", accessToken)
		params := []multipart.Param{
			multipart.Param{"img", fileName, fp},
		}

		b := &bytes.Buffer{}
		contentType, _ := multipart.Create(b, "", params)
		body = b.Bytes()
		headers = map[string]string{"Content-Type": contentType}
		return
	}

	return scanQrcode(cfgName, genParams)
}

func scanQrcode(cfgName string, genParams auth.FnGeneParams) (*QRCodeRes, error) {
	var res struct {
		callwxmp.BaseResult
		QRCodeRes
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res); err != nil {
		return nil, err
	}
	return &res.QRCodeRes, nil
}

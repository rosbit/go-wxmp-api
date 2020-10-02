// 本接口提供基于小程序的条码/二维码识别的API
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html

package img

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/img-util"
	"github.com/rosbit/go-wxmp-api/auth"
	"io"
)

type QRCodeRes struct {
	CodeResults []struct{
		TypeName string `json:"type_name"`
		Data string `json:"data"`
		Pos `json:"pos"`
	} `json:"code_results"`
	ImgSize `json:"img_size"`
}

func ScanQrcodeImgUrl(cfgName, imgUrl string) (*QRCodeRes, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = imgutil.GenerateUrl("https://api.weixin.qq.com/cv/img/qrcode", accessToken, imgUrl)
		return
	}

	return scanQrcode(cfgName, genParams)
}

func ScanQrcodeImg(cfgName, fileName string, fp io.Reader) (*QRCodeRes, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		return imgutil.GenerateImgMultipartParams(
			"https://api.weixin.qq.com/cv/img/qrcode",
			accessToken,
			[]multipart.Param{
				multipart.Param{"img", fileName, fp},
			},
		)
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

// 本接口提供基于小程序的营业执照 OCR 识别
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.businessLicense.html

package ocr

import (
	"github.com/rosbit/multipart-creator"
	"github.com/rosbit/go-wxmp-api/call-wxmp"
	"github.com/rosbit/go-wxmp-api/auth"
	"github.com/rosbit/go-wxmp-api/img-util"
	"io"
)

type BusinessLicense struct {
	RegisterNo string `json:"reg_num"` // 注册号
	Serial     string `json:"serial"`  // 编号
	LegalRepresentative string `json:"legal_representative"` // 法定代表人姓名
	EnterpriseName      string `json:"enterprise_name"`      // 企业名称
	TypeOfOrganization  string `json:"type_of_organization"` // 组成形式
	Address             string `json:"address"`            // 经营场所/企业住所
	TypeOfEnterprise    string `json:"type_of_enterprise"` // 公司类型
	BusinessScope       string `json:"business_scope"`     // 经营范围
	RegisteredCapital   string `json:"registered_capital"` // 注册资本
	PaidInCapital  string `json:"paid_in_capital"` // 实收资本
	ValidPeriod    string `json:"valid_period"`    // 营业期限
	RegisteredDate string `json:"registered_date"` // 注册日期/成立日期
	CertPosition struct {
		Pos PosT `json:"pos"`
	} `json:"cert_position"`  // 营业执照位置
	ImgSize ImgSizeT `json:"img_size"` // 图片大小
}

func ScanBusinessLicenseImgUrl(cfgName, imgUrl string) (*BusinessLicense, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		url = imgutil.GenerateUrl("https://api.weixin.qq.com/cv/ocr/bizlicense", accessToken, imgUrl)
		return
	}

	return scanBusinessLicense(cfgName, genParams)
}

func ScanBusinessLicenseImg(cfgName, fileName string, fp io.Reader) (*BusinessLicense, error) {
	genParams := func(accessToken string) (url string, body interface{}, headers map[string]string) {
		return imgutil.GenerateImgMultipartParams(
			"https://api.weixin.qq.com/cv/ocr/bizlicense",
			accessToken,
			[]multipart.Param{
				multipart.Param{"img", fileName, fp},
			},
		)
	}

	return scanBusinessLicense(cfgName, genParams)
}

func scanBusinessLicense(cfgName string, genParams auth.FnGeneParams) (*BusinessLicense, error) {
	var res struct {
		callwxmp.BaseResult
		BusinessLicense
	}
	if err := auth.CallWxmp(cfgName, genParams, "POST", callwxmp.HttpCall, &res); err != nil {
		return nil, err
	}
	return &res.BusinessLicense, nil
}

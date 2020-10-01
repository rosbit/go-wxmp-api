package imgutil

import (
	"github.com/rosbit/multipart-creator"
	"bytes"
	"fmt"
	"strings"
)

func GenerateImgMultipartParams(apiBaseUrl, accessToken string, params []multipart.Param) (
	url string,
	body interface{},
	headers map[string]string,
) {
	deli := '?'
	if strings.IndexByte(apiBaseUrl, '?') >= 0 {
		deli = '&'
	}
	url = fmt.Sprintf("%s%caccess_token=%s", apiBaseUrl, deli, accessToken)

	b := &bytes.Buffer{}
	contentType, _ := multipart.Create(b, "", params)
	body = b

	headers = map[string]string{"Content-Type": contentType}
	return
}

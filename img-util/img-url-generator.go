package imgutil

import (
	"net/url"
	"fmt"
	"strings"
)

func GenerateUrl(apiBaseUrl, accessToken, imgUrl string) string {
	v := url.Values{}
	v.Set("img_url", imgUrl)
	v.Set("access_token", accessToken)
	deli := '?'
	if strings.IndexByte(apiBaseUrl, '?') >= 0 {
		deli = '&'
	}
	return fmt.Sprintf("%s%c%s", apiBaseUrl, deli, v.Encode())
}

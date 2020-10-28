// 开发数据校验与解密
// 1. 对称解密使用的算法为 AES-128-CBC，数据采用PKCS#7填充。
// 2. 对称解密的目标密文为 Base64_Decode(encryptedData)。
// 3. 对称解密秘钥 aeskey = Base64_Decode(session_key), aeskey 是16字节。
// 4. 对称解密算法初始向量 为Base64_Decode(iv)，其中iv由数据接口返回。
// 参考文档: https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html

package decrypt

import (
	"encoding/base64"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
)

var (
	ErrInvalidBlockSize = fmt.Errorf("invalid block size")
	ErrInvalidPKCS7Data = fmt.Errorf("invalid PKCS7 data")
	ErrInvalidPKCS7Padding = fmt.Errorf("invalid padding on input")
)

// 数据解密
// @param sessionKey     auth.code2session获得的sessionKey
// @param encryptedData  小程序前端获得的加密串
// @param iv             小程序前端获得的解密算法初始向量
func Decrypt(sessionKey, encryptedData, iv string) ([]byte, error) {
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(plainText, cipherText)
	res, err := pkcs7Unpad(plainText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}

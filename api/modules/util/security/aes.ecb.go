package security

import (
	"crypto/aes"
)

//AesECBDecrypt AesECBDecrypt
func AesECBDecrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	return cbcDecrypt(key, ciphertext)
}

//unpadding
func unPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
}

func cbcDecrypt(key, src []byte) ([]byte, error) {
	//key只能是 16 24 32长度
	//返回加密结果
	dst := make([]byte, len(src))

	cryptBlocks(key, dst, src)
	return unPadding(dst), nil
}
func cryptBlocks(key, dst, src []byte) {
	blockSize := aes.BlockSize
	if len(src)%blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}

	//key只能是 16 24 32长度
	block, _ := aes.NewCipher(key)

	for len(src) > 0 {
		block.Decrypt(dst, src[:blockSize])
		src = src[blockSize:]
		dst = dst[blockSize:]
	}
	return
}

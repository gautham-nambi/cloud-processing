package hasher

import (
	"crypto/aes"
	"crypto/cipher"
)

func (i *Input) VerifyCipheredText() bool {
	key := []byte(createHash(Passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	data := []byte(i.CipheredText)
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext) == i.CipheredText
}

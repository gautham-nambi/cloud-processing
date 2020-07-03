package hasher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func (i *Input) Encrypt() {
	block, _ := aes.NewCipher([]byte(createHash(Passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	i.CipheredText = string(gcm.Seal(nonce, nonce, []byte(i.Text), nil))
}

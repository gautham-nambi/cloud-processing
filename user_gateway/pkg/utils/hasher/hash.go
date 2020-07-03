package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type Input struct {
	Text         string
	CipheredText string
}

const Passphrase = "SECRET"

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

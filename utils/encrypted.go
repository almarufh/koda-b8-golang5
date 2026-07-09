package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Encrypted(password string) string {
	encrypt := md5.Sum([]byte(password))
	return hex.EncodeToString(encrypt[:])
}

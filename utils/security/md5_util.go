package security

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data , salt []byte) string {
	md5 := md5.New()
	md5.Write(data)
	bytes := md5.Sum(salt)
	return hex.EncodeToString(bytes)
}

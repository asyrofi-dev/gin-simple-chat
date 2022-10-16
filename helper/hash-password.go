package helper

import (
	"crypto/sha1"
	"fmt"
)

func HashPassword(password *string) {
	sha := sha1.New()
	sha.Write([]byte(*password))
	encrypted := sha.Sum(nil)
	*password = fmt.Sprintf("%x", encrypted)
}

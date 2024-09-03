package utils

import (
	"crypto/rand"
	"unsafe"
)

func GenerateRandomString(str string, size int) string {
	s := []byte(str)
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = s[b[i]%byte(len(s))]
	}
	return *(*string)(unsafe.Pointer(&b))
}

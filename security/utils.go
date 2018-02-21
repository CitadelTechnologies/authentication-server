package security

import(
    "fmt"
    "crypto/rand"
)

func GenerateRandomToken(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return []byte(fmt.Sprintf("%x", b))
}

package model

import (
	"crypto/rand"
	"fmt"
)

func generateHash() string {
	b := make([]byte, 2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

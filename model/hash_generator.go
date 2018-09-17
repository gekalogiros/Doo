package model

import (
	"crypto/rand"
	"fmt"
)

func generateHash() string {
	b := make([]byte, 3)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
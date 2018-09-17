package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	hash := generateHash()
	assert.Equal(t, 6, len(hash), "")
}

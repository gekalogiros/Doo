package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHash(t *testing.T) {
	hash := generateHash()
	assert.Equal(t, 4, len(hash), "")
}

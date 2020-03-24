package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFeature(t *testing.T) {
	f, err := NewFeature("Fqy")
	assert.NotNil(t, f)
	assert.Nil(t, err)
}
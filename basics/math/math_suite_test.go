package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddnumbers(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result, "Addition result incorrect")
}

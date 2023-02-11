package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadData(t *testing.T) {
	strs, err := readData("states.json")
	assert.NoError(t, err)
	assert.Equal(t, 44, len(strs))
	assert.Contains(t, strs[0], "Washington")
}

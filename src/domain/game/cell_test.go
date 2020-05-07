package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCellConstants(t *testing.T) {
	assert.EqualValues(t, "FLAG", FLAG)
	assert.EqualValues(t, "QUESTION", QUESTION)
	assert.EqualValues(t, "NONE", NONE)
}

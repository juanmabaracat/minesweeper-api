package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCellConstants(t *testing.T) {
	assert.EqualValues(t, "flag", FLAG)
	assert.EqualValues(t, "question", QUESTION)
	assert.EqualValues(t, "none", NONE)
}

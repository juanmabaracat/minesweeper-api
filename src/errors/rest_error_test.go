package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("error message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "error message", err.Message())
	assert.EqualValues(t, "bad_request", err.Error())
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("test message", errors.New("database error"))
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "internal_server_error", err.Error())
	assert.EqualValues(t, "test message", err.Message())

	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

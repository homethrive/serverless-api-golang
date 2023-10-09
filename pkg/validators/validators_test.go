package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsEmailValid calls validators.IsEmailValid with an email, checking
// for a valid return value.
func TestIsEmailValid(t *testing.T) {
	email := "bob@mail.com"
	resp := IsEmailValid(email)
	assert.Equal(t, true, resp)
}

// TestIsEmailValidEmpty calls validators.IsEmailValid with an empty string,
// checking for an error.
func TestIsEmailValidEmpty(t *testing.T) {
	resp := IsEmailValid("")
	assert.Equal(t, false, resp)
}

// TestIsEmailValidInvalid calls validators.IsEmailValid with an empty string,
// checking for an error.
func TestIsEmailValidInvalid(t *testing.T) {
	resp := IsEmailValid("bobmail.com")
	assert.Equal(t, false, resp)
}

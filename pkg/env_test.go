package pkg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	expected := "gopher"
	os.Setenv("TEST_ENV", expected)
	result := GetEnv("TEST_ENV", "")

	assert.Equal(t, expected, result, "They should be equal")
}

func TestDefaultValue(t *testing.T) {
	expected := "gopher"
	result := GetEnv("TEST", expected)

	assert.Equal(t, expected, result, "They should be equal")
}

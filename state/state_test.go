package state

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewFromFile(t *testing.T) {
	// Arrange.
	_, err := NewFromFile("testdata/config.json")

	// Assert.
	require.Nil(t, err)
}

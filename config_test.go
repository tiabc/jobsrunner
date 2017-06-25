package jobsrunner

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConfigJobInterval_UnmarshalJSON_Null(t *testing.T) {
	// Arrange.
	var c ConfigJobInterval

	// Act.
	err := c.UnmarshalJSON([]byte("null"))

	// Assert.
	require.Nil(t, err)
	require.EqualValues(t, 0, c)
}

func TestConfigJobInterval_UnmarshalJSON_TooManySpaces(t *testing.T) {
	// Arrange.
	var c ConfigJobInterval

	// Act.
	err := c.UnmarshalJSON([]byte("5  seconds"))

	// Assert.
	require.NotNil(t, err)
}

func TestConfigJobInterval_UnmarshalJSON_CannotParseNumber(t *testing.T) {
	// Arrange.
	var c ConfigJobInterval

	// Act.
	err := c.UnmarshalJSON([]byte("five seconds"))

	// Assert.
	require.NotNil(t, err)
}

func TestConfigJobInterval_UnmarshalJSON_NonpositiveNumber(t *testing.T) {
	// Arrange.
	var c ConfigJobInterval

	// Act.
	err1 := c.UnmarshalJSON([]byte("-3 seconds"))
	err2 := c.UnmarshalJSON([]byte("0 seconds"))

	// Assert.
	require.NotNil(t, err1)
	require.NotNil(t, err2)
}

func TestConfigJobInterval_UnmarshalJSON_CorrectParsing(t *testing.T) {
	// Arrange.
	var c ConfigJobInterval

	t.Run("second", func(t *testing.T) {
		// Act.
		err := c.UnmarshalJSON([]byte("1 second"))

		// Assert.
		require.Nil(t, err)
		require.EqualValues(t, 1*time.Second, c)
	})

	t.Run("seconds", func(t *testing.T) {
		// Act.
		err := c.UnmarshalJSON([]byte("3 seconds"))

		// Assert.
		require.Nil(t, err)
		require.EqualValues(t, 3*time.Second, c)
	})

	t.Run("minute", func(t *testing.T) {
		// Act.
		err := c.UnmarshalJSON([]byte("1 minute"))

		// Assert.
		require.Nil(t, err)
		require.EqualValues(t, 1*time.Minute, c)
	})

	t.Run("hours", func(t *testing.T) {
		// Act.
		err := c.UnmarshalJSON([]byte("3 hours"))

		// Assert.
		require.Nil(t, err)
		require.EqualValues(t, 3*time.Hour, c)
	})

	t.Run("hour", func(t *testing.T) {
		// Act.
		err := c.UnmarshalJSON([]byte("1 hour"))

		// Assert.
		require.Nil(t, err)
		require.EqualValues(t, 1*time.Hour, c)
	})
}

func TestConfigJobInterval_UnmarshalJSON_BadModifier(t *testing.T) {
	// Arrange.
	var c ConfigJobInterval

	// Act.
	err := c.UnmarshalJSON([]byte("5 daysOrSo"))

	// Assert.
	require.NotNil(t, err)
}

func TestNewConfigFromFile(t *testing.T) {
	// Arrange.
	_, err := NewConfigFromFile("absolutely-absent-file.json")

	// Assert.
	require.NotNil(t, err)
}

func TestNewConfigFromFile_(t *testing.T) {
	// Arrange.
	c, err := NewConfigFromFile("testdata/config.json")

	// Assert.
	require.Nil(t, err)
	require.EqualValues(t, 1, c.Version)
	require.Len(t, c.Jobs, 1)
	require.EqualValues(t, "yourapp check-statuses", c.Jobs[0].Cmd)
	require.EqualValues(t, 5*time.Second, c.Jobs[0].Interval)
}

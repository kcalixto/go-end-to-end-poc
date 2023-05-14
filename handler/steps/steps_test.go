package steps

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSuccess(t *testing.T) {
	// Test 1
	res, err := runStep["request"].Execute(guy)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	res, err = runStep["provide"].Execute(sparkles)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	res, err = runStep["get"].Execute(guy)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestError(t *testing.T) {
	// Test 2
	res, err = runStep["request"].Execute(otherGuy)
	require.Error(t, err)
	require.Empty(t, res)

	res, err = runStep["provide"].Execute(darkness)
	require.Error(t, err)
	require.Empty(t, res)

	res, err = runStep["get"].Execute(otherGuy)
	require.Error(t, err)
	require.Empty(t, res)
}

func TestGuyWithoutPinkSparkle(t *testing.T) {
	// Test 3
	res, err = runStep["request"].Execute(guy)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	res, err = runStep["provide"].Execute(pinky)
	require.NoError(t, err)
	require.Contains(t, res, "Sparkle")
}

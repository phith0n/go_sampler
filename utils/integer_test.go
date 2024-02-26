package utils

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestEnsureUInt(t *testing.T) {
	require.Equal(t, uint(0), EnsureUInt(-1))
	require.Equal(t, uint(0), EnsureUInt("-1"))
	require.Equal(t, uint(0), EnsureUInt(0))
	require.Equal(t, uint(123), EnsureUInt(123))
	require.Equal(t, uint(2251799813685248), EnsureUInt("2251799813685248")) // 2<50
	require.Equal(t, uint(0), EnsureUInt("2417851639229258349412352"))       // 2<<80, over buffer
	require.Equal(t, uint(0), EnsureUInt(1.2))
}

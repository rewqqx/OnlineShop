package utils

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHashPassword(t *testing.T) {
	newHashPassword := HashPassword("test")

	expectedHashPassword := "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"

	require.Equal(t, newHashPassword, expectedHashPassword, "Invalid hash of "+
		"the password. Expected %v, got %v", expectedHashPassword, newHashPassword)
}

func TestGenerateToken(t *testing.T) {
	tokenLength := 16
	token := GenerateToken(tokenLength)

	require.Equal(t, len(token), tokenLength*2, "Token length is %d, expected %d", len(token), tokenLength*2)

	decoded, err := hex.DecodeString(token)
	if err != nil {
		t.Error("Error decoding token:", err)
	}
	require.Equal(t, len(decoded), tokenLength, "Invalid decoding token. Expected %d, actually %d",
		tokenLength, len(decoded))

	for i := 0; i < 1000; i++ {
		newGenerateToken := GenerateToken(tokenLength)
		require.NotEqual(t, newGenerateToken, token, "Repeating identical tokens")
	}

}

package utils

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"os"
	"reflect"
	"testing"
	"time"
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

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	timestamp := Timestamp{Time: time.Now(), Valid: true}

	data := []byte("2006-01-02 15:04:05")
	err := timestamp.UnmarshalJSON(data)

	require.Equal(t, err, nil, "Expected that func UnmarshalJSON() will not return an error")
}

func TestTimestamp_UnmarshalJSONWithError(t *testing.T) {
	timestamp := Timestamp{Time: time.Now(), Valid: true}

	invalidData := []byte("\"\\x01\"")
	err := timestamp.UnmarshalJSON(invalidData)
	if err == nil {
		t.Error("Expected that func UnmarshalJSON() will return an error")
	}
}

func TestTimestamp_MarshalJSONWithValidTime(t *testing.T) {
	timestamp := &Timestamp{Time: time.Now(), Valid: true}

	JSON, err := timestamp.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error in MarshalJSON(): %v", err)
	}

	require.Equal(t, string(JSON), time.Now().Format("\"2006-01-02 15:04:05\""))
}

func TestTimestamp_MarshalJSONWithUnValidTime(t *testing.T) {
	timestamp := &Timestamp{Time: time.Now(), Valid: false}

	JSON, err := timestamp.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error in MarshalJSON(): %v", err)
	}

	require.Equal(t, string(JSON), "")
}

func TestTimestamp_Value(t *testing.T) {
	timestamp := Timestamp{Time: time.Now(), Valid: true}

	value, err := timestamp.Value()
	if err != nil {
		t.Error("Return unexpected error in Value():", err)
	}

	actuallyValueType := reflect.TypeOf(value)
	expectedType := reflect.TypeOf(time.Time{})
	require.Equal(t, actuallyValueType, expectedType, "Value expected %s, actually %s", expectedType, actuallyValueType)
}

func TestTimestamp_Scan(t *testing.T) {
	sampleTimestamp := Timestamp{Valid: false}

	inputTime := time.Date(2023, 4, 12, 14, 30, 0, 0, time.UTC)

	err := sampleTimestamp.Scan(inputTime)
	require.Equal(t, err, nil, "Returned unexpected err: %s", err)
}

func TestOpen(t *testing.T) {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	dbConnection := &DBConnect{
		Ip:       host,
		Port:     "5432",
		User:     "postgres",
		Password: "pgpass",
		Database: "postgres",
	}

	go func() {
		err := dbConnection.Open()
		require.Equal(t, err, nil)

		err = dbConnection.Connection.Close()
		require.Equal(t, err, nil)
	}()
}

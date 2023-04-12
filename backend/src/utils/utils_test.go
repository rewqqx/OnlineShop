package utils

import (
	"encoding/hex"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
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
	require.Equal(t, err, nil, "Error decoding token: %v", err)
	require.Equal(t, len(decoded), tokenLength, "Invalid decoding token. Expected %d, actually %d",
		tokenLength, len(decoded))

	for i := 0; i < 1000; i++ {
		newGenerateToken := GenerateToken(tokenLength)
		require.NotEqual(t, newGenerateToken, token, "Repeating identical tokens")
	}

}

func TestTimestampUnmarshalJSON(t *testing.T) {
	timestamp := Timestamp{Time: time.Now(), Valid: true}

	data := []byte("2006-01-02 15:04:05")
	err := timestamp.UnmarshalJSON(data)

	require.Equal(t, err, nil, "Expected that func UnmarshalJSON() will not return an error")
}

func TestTimestampUnmarshalJSONWithError(t *testing.T) {
	timestamp := Timestamp{Time: time.Now(), Valid: true}

	invalidData := []byte("\"\\x01\"")
	err := timestamp.UnmarshalJSON(invalidData)
	if err == nil {
		t.Error("Expected that func UnmarshalJSON() will return an error")
	}
}

func TestTimestampMarshalJSONWithValidTime(t *testing.T) {
	timestamp := &Timestamp{Time: time.Now(), Valid: true}

	JSON, err := timestamp.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error in MarshalJSON(): %v", err)
	}

	require.Equal(t, string(JSON), time.Now().Format("\"2006-01-02 15:04:05\""))
}

func TestTimestampMarshalJSONWithUnValidTime(t *testing.T) {
	timestamp := &Timestamp{Time: time.Now(), Valid: false}

	JSON, err := timestamp.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error in MarshalJSON(): %v", err)
	}

	require.Equal(t, string(JSON), "")
}

func TestTimestampValue(t *testing.T) {
	timestamp := Timestamp{Time: time.Now(), Valid: true}

	value, err := timestamp.Value()
	if err != nil {
		t.Error("Return unexpected error in Value():", err)
	}

	actuallyValueType := reflect.TypeOf(value)
	expectedType := reflect.TypeOf(time.Time{})
	require.Equal(t, actuallyValueType, expectedType, "Value expected %s, actually %s", expectedType, actuallyValueType)
}

func TestTimestampScan(t *testing.T) {
	sampleTimestamp := Timestamp{Valid: false}

	inputTime := time.Date(2023, 4, 12, 14, 30, 0, 0, time.UTC)

	err := sampleTimestamp.Scan(inputTime)
	require.Equal(t, err, nil, "Returned unexpected err: %s", err)
}

func TestDBConnectGetTables(t *testing.T) {
	client := &DBConnect{
		Ip:       "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "pgpass",
		Database: "postgres",
	}

	db, mock, err := sqlmock.New()
	require.Equal(t, err, nil, "Failed of creating DB: %v", err)

	client.Connection = sqlx.NewDb(db, "postgres")

	rows := sqlmock.NewRows([]string{"table_name"}).AddRow("table1").AddRow("table2")
	mock.ExpectQuery("SELECT table_name FROM information_schema.tables").WithArgs("test_schema").WillReturnRows(rows)

	tables, err := client.GetTables("test_schema")
	require.Equal(t, err, nil, "Unexpected error: %v", err)
	require.Equal(t, len(tables), 2, "Expected %d, actually %d raws in tables", 2, len(tables))
	require.EqualValues(t, tables, []string{"table1", "table2"}, "Expected %s, got %s",
		[]string{"table1", "table2"}, tables)
}

func TestDBConnectGetSchemas(t *testing.T) {
	client := &DBConnect{
		Ip:       "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "pgpass",
		Database: "postgres",
	}

	db, mock, err := sqlmock.New()
	require.Equal(t, err, nil, "Failed of creating DB: %v", err)

	client.Connection = sqlx.NewDb(db, "postgres")

	rows := sqlmock.NewRows([]string{"schema_name"}).AddRow("public").AddRow("test_schema").AddRow("testTwo")
	mock.ExpectQuery("SELECT schema_name FROM information_schema.schemata").WillReturnRows(rows)

	schemas, err := client.GetSchemas()
	require.Equal(t, err, nil, "Error in GetSchemas(): %v", err)
	require.Equal(t, len(schemas), 3, "Expected %d, got %d rows in schemas", 3, len(schemas))
	require.EqualValues(t, schemas, []string{"public", "test_schema", "testTwo"}, "Expected %s, got %s",
		[]string{"public", "test_schema", "testTwo"}, schemas)
}

package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisPing(t *testing.T) {
	rds := New("127.0.0.1", "6379", "")
	err := rds.Ping()
	assert.Equal(t, nil, err)
}

package adapter

import (
	"backend/src/utils/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getRedisClient(t *testing.T) *ItemCartDatabase {
	rds := database.New("127.0.0.1", "6379", "")
	err := rds.Ping()
	assert.Equal(t, nil, err)
	return &ItemCartDatabase{rds}
}

func TestAddItem(t *testing.T) {
	adapter := getRedisClient(t)
	err := adapter.AddItem(CartItem{ItemID: 0, UserID: 0, Count: 10})
	assert.Equal(t, nil, err)
}

func TestGetCart(t *testing.T) {
	adapter := getRedisClient(t)
	err := adapter.AddItem(CartItem{ItemID: 0, UserID: 0, Count: 10})
	assert.Equal(t, nil, err)
	err = adapter.AddItem(CartItem{ItemID: 1, UserID: 0, Count: 3})
	assert.Equal(t, nil, err)
	items, err := adapter.GetCart(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, len(items) >= 2)
}

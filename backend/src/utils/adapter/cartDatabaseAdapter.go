package adapter

import (
	"backend/src/utils/database"
	"fmt"
	"time"
)

type ItemCartDatabase struct {
	redis *database.Redis
}
type CartItem struct {
	ItemID int `json:"item_id" db:"item_id"`
	UserID int `json:"user_id" db:"user_id"`
	Count  int `json:"count" db:"count"`
}

func (item *CartItem) getRedisKey() string {
	return fmt.Sprintf("cart::%v", item.UserID)
}

func (item *CartItem) getItemKey() string {
	return fmt.Sprintf("%v", item.ItemID)
}

func (adapter *ItemCartDatabase) AddItem(item *CartItem) error {
	key := item.getRedisKey()
	itemKey := item.getItemKey()

	pipe := adapter.redis.Client.Pipeline()
	cmd := pipe.HMGet(key, itemKey)

	curCount := 0

	if cmd.Val() != nil {
		curCount = cmd.Val()[0].(int)
	}

	input := make(map[string]interface{})
	input[itemKey] = curCount + item.Count

	pipe.HMSet(key, input)
	_, err := pipe.Exec()

	if err != nil {
		return cmd.Err()
	}

	boolCmd := adapter.redis.Client.Expire(key, 24*7*time.Hour)
	return boolCmd.Err()
}

func (adapter *ItemCartDatabase) DeleteItem(item *CartItem) error {
	key := item.getRedisKey()
	itemKey := item.getItemKey()

	cmd := adapter.redis.Client.HDel(key, itemKey)

	if cmd.Err() != nil {
		return cmd.Err()
	}

	boolCmd := adapter.redis.Client.Expire(key, 24*7*time.Hour)
	return boolCmd.Err()
}

func (adapter *ItemCartDatabase) SetItem(item *CartItem) error {
	key := item.getRedisKey()
	itemKey := item.getItemKey()

	input := make(map[string]interface{})
	input[itemKey] = item.Count

	cmd := adapter.redis.Client.HMSet(key, input)

	if cmd.Err() != nil {
		return cmd.Err()
	}

	boolCmd := adapter.redis.Client.Expire(key, 24*7*time.Hour)
	return boolCmd.Err()
}

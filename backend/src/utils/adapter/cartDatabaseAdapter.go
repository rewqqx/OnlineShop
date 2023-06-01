package adapter

import (
	"backend/src/utils/database"
	"fmt"
	"strconv"
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

func CreateCartDatabaseAdapter(redis *database.Redis) *ItemCartDatabase {
	adapter := &ItemCartDatabase{redis: redis}
	return adapter
}

func (item *CartItem) getRedisKey() string {
	return fmt.Sprintf("cart::%v", item.UserID)
}

func (item *CartItem) getItemKey() string {
	return fmt.Sprintf("%v", item.ItemID)
}

func (adapter *ItemCartDatabase) AddItem(item CartItem) error {
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

func (adapter *ItemCartDatabase) DeleteItem(item CartItem) error {
	key := item.getRedisKey()
	itemKey := item.getItemKey()

	cmd := adapter.redis.Client.HDel(key, itemKey)

	if cmd.Err() != nil {
		return cmd.Err()
	}

	boolCmd := adapter.redis.Client.Expire(key, 24*7*time.Hour)
	return boolCmd.Err()
}

func (adapter *ItemCartDatabase) DeleteItems(item CartItem) error {
	key := item.getRedisKey()

	cmd := adapter.redis.Client.Del(key)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return cmd.Err()
}

func (adapter *ItemCartDatabase) SetItem(item CartItem) error {
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

func (adapter *ItemCartDatabase) GetCart(userID int) (output []CartItem, err error) {
	key := fmt.Sprintf("cart::%v", userID)

	cmd := adapter.redis.Client.HGetAll(key)

	result, err := cmd.Result()

	if err != nil {
		return
	}

	for k, v := range result {
		itemID, subErr := strconv.Atoi(k)
		if err != nil {
			return output, subErr
		}
		count, subErr := strconv.Atoi(v)
		if err != nil {
			return output, subErr
		}
		output = append(output, CartItem{itemID, userID, count})
	}

	return
}

package adapter

import (
	"backend/src/utils/database"
	"fmt"
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
	input[itemKey] = curCount

	pipe.HMSet(key, input)
	_, err := pipe.Exec()
	return err
}

func (adapter *ItemCartDatabase) DeleteItem(item *CartItem) {

}

func (adapter *ItemCartDatabase) SetItem(item *CartItem) {

}

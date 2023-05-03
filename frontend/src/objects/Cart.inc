<?php

class Cart
{
    private mixed $user_id = 0;

    private mixed $items = array();

    public function  __construct($data){
        $this->user_id = $data["id"] ?: 0;
        $this->items = $data['items'] ?: array();
    }

    /**
     * @return int|mixed
     */
    public function getUserId(): mixed
    {
        return $this->user_id;
    }

    /**
     * @return array|mixed
     */

    # [ 0 => ['item_id' => 1, 'price' => 50, 'quantity' => 2]]
    public function getItems(): mixed
    {
//        return $this->items;
        return [['item_id' => 1, 'price' => 50, 'quantity' => 2]];
    }


    public function toJson(): string
    {
        $array["user_id"] = $this->user_id;
        $array["items"] = $this->items;

        return json_encode($array);
    }




}
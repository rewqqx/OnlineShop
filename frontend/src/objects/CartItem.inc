<?php

class CartItem
{
    private $id = -1;
    private $userID = -1;
    private $count = 0;

    public function __construct($data)
    {
        $this->id = $data["item_id"] ?: -1;
        $this->userID = $data["user_id"] ?: -1;
        $this->count = $data["count"] ?: 0;
    }


    public function getId(): int
    {
        return $this->id;
    }


    public function getUserID(): string
    {
        return $this->userID;
    }

    public function getCount(): array
    {
        return $this->count;
    }

    public function toJson(): string
    {
        $array["item_id"] = $this->id;
        $array["user_id"] = $this->userID;
        $array["count"] = $this->count;

        return json_encode($array);
    }


}

?>
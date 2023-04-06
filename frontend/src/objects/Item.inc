<?php

class Item
{
    private $id = -1;
    private $name = "Default";
    private $price = 0;
    private $description = "";
    private $imageIDs = array();

    public function __construct($data)
    {
        $this->id = $data["id"] ?: -1;
        $this->name = $data["name"] ?: "Default";
        $this->price = $data["price"] ?: 0;
        $this->description = $data["description"] ?: "";
        $this->imageIDs = $data["image_ids"] ?: array();
    }


    public function getId(): int
    {
        return $this->id;
    }

    public function getName(): string
    {
        return $this->name;
    }


    public function getPrice(): int
    {
        return $this->price;
    }


    public function getDescription(): string
    {
        return $this->description;
    }

    public function getImageIDs(): array
    {
        return $this->imageIDs;
    }

    public function toJson(): string
    {
        $array["id"] = $this->id;
        $array["name"] = $this->name;
        $array["price"] = $this->price;
        $array["description"] = $this->description;
        $array["image_ids"] = $this->imageIDs;

        return json_encode($array);
    }


}

?>
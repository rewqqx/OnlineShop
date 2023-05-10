<?php

class Tag
{
    private $id = -1;
    private $name = "Default";
    private $parentID = -1;

    public function __construct($data)
    {
        $this->id = $data["id"] ?: -1;
        $this->name = $data["tag_name"] ?: "Default";
        $this->parentID = $data["parent_id"] ?: -1;
    }


    public function getId(): int
    {
        return $this->id;
    }

    public function getName(): string
    {
        return $this->name;
    }



    public function toJson(): string
    {
        $array["id"] = $this->id;
        $array["tag_name"] = $this->name;
        $array["parent_id"] = $this->parentID;

        return json_encode($array);
    }


}

?>
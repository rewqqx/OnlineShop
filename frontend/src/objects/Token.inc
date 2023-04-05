<?php

class Token
{

    private $id;
    private $token;


    function __construct($data)
    {
        $this->id = $data["ID"];
        $this->token = $data["token"];
    }

    public function toJson(): string
    {
        $array["ID"] = $this->id;
        $array["token"] = $this->token;
        return json_encode($array);
    }

    function getID(): int
    {
        return $this->id;
    }

    function getToken(): string
    {
        return $this->token;
    }

}

?>
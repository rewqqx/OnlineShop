<?php

class Token
{

    public $id;
    public $token;


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


}

?>
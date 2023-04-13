<?php

class Token
{

    private $id = "";
    private $token = "";

    private $valid = false;


    function __construct($data)
    {
        if (array_key_exists("ID", $data)) {
            $this->id = $data["ID"];
        }

        if (array_key_exists("token", $data)) {
            $this->token = $data["token"];
            $this->valid = true;
        }
    }

    public function isValid(): bool
    {
        if (!$this->valid) {
            return false;
        }

        $user = GetUser($this);

        return $user != null;
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
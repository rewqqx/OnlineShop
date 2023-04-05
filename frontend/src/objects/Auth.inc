<?php

class Auth
{

    public $mail;
    public $password;


    function __construct($mail, $password)
    {
        $this->password = $password;
        $this->mail = $mail;
    }


    public function toJson(): string
    {
        $array["mail"] = $this->mail;
        $array["password"] = $this->password;
        return json_encode($array);
    }


}

?>
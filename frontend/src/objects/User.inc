<?php

class User
{
    private $id;
    private $user_name;
    private $user_surname;
    private $user_patronymic;
    private $phone;
    private $birthdate;
    private $password_hash;
    private $mail;
    private $role_id;

    private $token;


    function __construct($data)
    {
        $this->id = $data['id'] ?? "";
        $this->user_name = $data['name'] ?? "";
        $this->user_surname = $data['user_surname'] ?? "";
        $this->user_patronymic = $data['user_patronymic'] ?? "";
        $this->phone = $data['phone'] ?? "";
        $this->birthdate = $data['birthdate'] ?? "";
        $this->password_hash = $data['password_hash'] ?? "";
        $this->mail = $data['mail'] ?? "";
        $this->token = $data['token'] ?? "";
        $this->role_id = $data['role_id'] ?? "";
    }

    /**
     * @return mixed|string
     */
    public function getId(): mixed
    {
        return $this->id;
    }

    /**
     * @return mixed|string
     */
    public function getUserName(): mixed
    {
        return $this->user_name ?? $this->id;
    }

    /**
     * @return mixed|string
     */
    public function getUserSurname(): mixed
    {
        return $this->user_surname;
    }

    /**
     * @return mixed|string
     */
    public function getUserPatronymic(): mixed
    {
        return $this->user_patronymic;
    }

    /**
     * @return mixed|string
     */
    public function getPhone(): mixed
    {
        return $this->phone;
    }


    /**
     * @return mixed|string
     */
    public function getBirthdate(): mixed
    {
        return $this->birthdate;
    }

    /**
     * @return mixed|string
     */
    public function getMail(): mixed
    {
        return $this->mail;
    }


    /**
     * @return mixed|string
     */
    public function getRoleId(): mixed
    {
        return $this->role_id;
    }


    public function toJson(): string
    {
        $array["id"] = $this->id;
        $array["user_name"] = $this->user_name;
        $array["user_surname"] = $this->user_surname;
        $array["user_patronymic"] = $this->user_patronymic;
        $array["phone"] = $this->phone;
        $array["birthday"] = $this->birthdate;
        $array["mail"] = $this->mail;
        $array["role_id"] = $this->role_id;
        $array["password_hash"] = $this->password_hash;
        $array["token"] = $this->token;

        return json_encode($array);
    }


}

?>
<?php

namespace OnlineShop\user;

class User
{
    public $id;
    public $user_name;
    public $user_surname;
    public $user_patronymic;
    public $phone;
    public $birthdate;
    public $password_hash;
    public $mail;
    public $role_id;


    function __construct($data){
        $this->id = $data['id'] ?? "";
        $this->user_name = $data['name'] ?? "";
        $this->user_surname = $data['user_surname'] ?? "";
        $this->user_patronymic = $data['user_patronymic'] ?? "";
        $this->phone = $data['phone'] ?? "";
        $this->birthdate = $data['birthdate'] ?? "";
        $this->password_hash = $data['password_hash'] ?? "";
        $this->mail = $data['mail'] ?? "";
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



    public function toJson($array) : string
    {
        return json_encode($array);
    }


}
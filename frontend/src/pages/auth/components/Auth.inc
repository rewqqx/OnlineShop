<?php

class Auth extends DOM
{

    private Text $title;
    private DOM $container;

    private Text $loginHeader;
    private Input $login;

    private Text $passwordHeader;
    private Input $password;

    private Button $confirm;

    /*
     * TODO: Очень много описаний стилей,
     * чтобы уменьшить это необходимо разработать
     * стайл код, чтобы максимально переиспользовать стили
     */
    protected function createElement()
    {
        $this->setStyle("auth-page");

        $this->title = new Text("Authentication");
        $this->title->setStyle("header-text")->setAttribute("color", "#2D2D2D");
        $this->title->setAttribute("text-align", "center")->setAttribute("padding", "10px");

        $this->container = new DOM();
        $this->container->setStyle("auth-container");

        $this->loginHeader = new Text("Login");
        $this->loginHeader->setStyle("button-text")->setAttribute("color", "#2D2D2D")->setAttribute("margin-left", "16px");

        $this->login = new Input();
        $this->login->setStyle("input")->setAttribute("margin-left", "15px")->setAttribute("margin-right", "15px");

        $this->passwordHeader = new Text("Password");
        $this->passwordHeader->setStyle("button-text")->setAttribute("color", "#2D2D2D")->setAttribute("margin-left", "16px");

        $this->password = new Input();
        $this->password->setStyle("input")->setAttribute("margin-left", "15px")->setAttribute("margin-right", "15px");

        $this->confirm = new Button("Sign in");
        $this->confirm->setAttribute("margin-left", "auto")->setAttribute("margin-right", "20px");
        $this->confirm->textElement->setAttribute("color", "#2D2D2D");

        $this->container->append($this->title)->append($this->loginHeader)->append($this->login);
        $this->container->append($this->passwordHeader)->append($this->password)->append($this->confirm);
        $this->append($this->container);

        return $this;
    }
}


?>
<?php

class AuthComponent extends DOM
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
        $this->setTag("form");
        $this->setStyle("auth-page");
        $this->setMethod("POST");

        $this->title = new Text("Login");
        $this->title->setStyle("header-text")->setAttribute("color", "#2E2E2E");
        $this->title->setAttribute("text-align", "center")->setAttribute("padding", "10px");

        $this->container = new DOM();
        $this->container->setStyle("auth-container");

        $this->loginHeader = new Text("LOGIN");
        $this->loginHeader->setStyle("input-header");

        $this->login = new Input("LOGIN");
        $this->login->setStyle("input")->setAttribute("margin-bottom", "20px");

        $this->passwordHeader = new Text("PASSWORD");
        $this->passwordHeader->setStyle("input-header");

        $this->password = new Input("PASSWORD");
        $this->password->setStyle("input")->setAttribute("margin-bottom", "20px");

        $parent = $this;


        $this->confirm = new Button("SIGN IN");
        $this->confirm->setAttribute("align-self", "center");
        $this->confirm->setAction(function () use ($parent) {
            $login = $_POST[$parent->login->name];
            $password = $_POST[$parent->password->name];
            $auth = new Auth($login, $password);
            $token = GetUserToken($auth);
            if ($token != null) {
                setTokenCookie($token);
                header('Location: /OnlineShop/frontend/src/index.php');
            }
        });


        $this->container->append($this->title)->append($this->loginHeader)->append($this->login);
        $this->container->append($this->passwordHeader)->append($this->password)->append($this->confirm);
        $this->append($this->container);

        return $this;
    }

}


?>
<?php

class CartComponent extends DOM
{
    private CartHeader $header;
    private Container $container;

    protected function createElement(): DOM
    {
        $this->header = new CartHeader();
        $this->container = new Container();
        $this->append($this->header)->append($this->container);

        $token = getTokenFromCookie();

        if (!$token->isValid()) {
            header('Location: /frontend/src/index.php');
        }

        $this->container->setTag("form");
        $this->container->setStyle("profile-page");
        $this->container->setMethod("POST");

        $user = GetUser($token);

        $panel = new DOM();
        $panel->setStyle("profile-panel");
        $this->container->append($panel);

        $nameHeader = new Text("NAME");
        $nameHeader->setStyle("input-header");
        $panel->append($nameHeader);

        $name = new Input("NAME");
        $name->setValue($user->getUserName());
        $name->setStyle("input")->setAttribute("margin-bottom", "20px");
        $panel->append($name);

        $surnameHeader = new Text("SURNAME");
        $surnameHeader->setStyle("input-header");
        $panel->append($surnameHeader);

        $surname = new Input("SURNAME");
        $surname->setValue($user->getUserSurname());
        $surname->setStyle("input")->setAttribute("margin-bottom", "20px");
        $panel->append($surname);

        $phoneHeader = new Text("PHONE");
        $phoneHeader->setStyle("input-header")->setAttribute("margin-top", "24px");;
        $panel->append($phoneHeader);

        $phone = new Input("PHONE");
        $phone->setValue($user->getPhone());
        $phone->setStyle("input")->setAttribute("margin-bottom", "20px");
        $panel->append($phone);

        $loginHeader = new Text("LOGIN");
        $loginHeader->setStyle("input-header")->setAttribute("margin-top", "24px");;
        $panel->append($loginHeader);

        $login = new Input("LOGIN");
        $login->setValue($user->getMail());
        $login->setStyle("input")->setAttribute("margin-bottom", "20px");
        $panel->append($login);

        $passwordHeader = new Text("PASSWORD");
        $passwordHeader->setStyle("input-header")->setAttribute("margin-top", "24px");;
        $panel->append($passwordHeader);

        $password = new Input("PASSWORD");
        $password->setStyle("input")->setAttribute("margin-bottom", "20px");
        $panel->append($password);


        $passwordConfirmHeader = new Text("PASSWORD CONFIRM");
        $passwordConfirmHeader->setStyle("input-header");
        $panel->append($passwordConfirmHeader);

        $confirmPassword = new Input("PASSWORD CONFIRM");
        $confirmPassword->setStyle("input")->setAttribute("margin-bottom", "20px");
        $panel->append($password);

        $parent = $this;

        $confirm = new Button("SAVE");
        $confirm->setAttribute("align-self", "center")->setAttribute("background", "#fafafa")->setAttribute("margin-top", "auto");
        $confirm->setAction(function () use ($parent) {
//                функцию чтобы добавлял item и количестов ссылка на рэдис
        });
        $panel->append($confirmPassword);
        return $this;
    }
}

<?php

class Header extends DOM
{

    private Link $profile;
    private Text $title;

    private IconButton $cart;

    protected function createElement(): DOM
    {
        $this->setTag("form");
        $this->title = new Text("Online shop");
        $this->title->setStyle("header-text");

        $token = getTokenFromCookie();

        $buttonName = "Sign In";
        $link = "/frontend/src/pages/auth/AuthPage.php";

        $this->cart = new IconButton("ic_cart_24x24.svg");
        $this->cart->setName("cart");
        $this->cart->setValue("1");
        $this->cart->setAttribute("margin-right", "10px");
        $this->cart->setAction(function ()  {
            header('Location: /frontend/src/pages/cart/CartPage.inc');
        });

        if ($token->isValid()) {
            $buttonName = "Profile";
            $link = "/frontend/src/pages/profile/ProfilePage.php";
        }

        $this->profile = new Link($buttonName);
        $this->profile->setAttribute("margin-left", "auto");

        $this->profile->setLink($link);

        $this->setStyle("header");

        $this->append($this->title);
        $this->append($this->profile);
        $this->append($this->cart);

        return $this;
    }
}

?>
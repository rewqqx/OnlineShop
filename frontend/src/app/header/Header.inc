<?php

class Header extends DOM
{

    private Link $profile;
    private Text $title;

    private IconButton $cart;

    protected function createElement(): DOM
    {
        $this->title = new Text("Online shop");
        $this->title->setStyle("header-text");

        $token = getTokenFromCookie();

        $buttonName = "Sign In";
        $link = "/OnlineShop/frontend/src/pages/auth/AuthPage.inc";

        $this->cart = new IconButton("ic_cart_24x24.svg");
        $this->cart->setAttribute("margin-right", "10px");

        if ($token->isValid()) {
            $buttonName = "Profile";
            $link = "/OnlineShop/frontend/src/pages/profile/ProfilePage.inc";
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
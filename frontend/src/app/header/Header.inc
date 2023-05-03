<?php

class Header extends DOM
{

    private Link $profile;
    private Text $title;

    private Link $cart;

    protected function createElement()
    {
        $this->title = new Text("Online shop");
        $this->title->setStyle("header-text");

        $token = getTokenFromCookie();

        $buttonName = "Sign In";
        $link = "/OnlineShop/frontend/src/pages/auth/AuthPage.inc";

        if ($token->isValid()) {
            $buttonName = "Profile";
            $link = "/OnlineShop/frontend/src/pages/profile/ProfilePage.inc";
            $buttonCart = "Cart";
            $linkCart = "/OnlineShop/frontend/src/pages/cart/CartPage.inc";

            $this->cart = new Link($buttonCart);
            $this->cart->setAttribute("margin-left", "1900px");
            $this->cart->setAttribute("margin-right", "40px");

            $this->cart->setLink($linkCart);

        }

        $this->profile = new Link($buttonName);
        $this->profile->setAttribute("margin-left", "auto");
        $this->profile->setAttribute("margin-right", "20px");

        $this->profile->setLink($link);

        $this->setStyle("header");

        $this->append($this->title);
        $this->append($this->cart);
        $this->append($this->profile);


        return $this;
    }
}

?>
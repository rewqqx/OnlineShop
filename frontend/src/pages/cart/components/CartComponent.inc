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
        $this->container->setStyle("cart-page");
        $this->container->setMethod("POST");

        $user = GetUser($token);

        $panelLeft = new DOM();
        $panelLeft->setStyle("cart-panel");
        $this->container->append($panelLeft);


        $panelRight = new DOM();
        $panelRight->setStyle("cart-filler");
        $this->container->append($panelRight);


        return $this;
    }
}

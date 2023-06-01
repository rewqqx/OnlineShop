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
        $cart = GetCartItems($user->getId());

        $panelLeft = new DOM();
        $panelLeft->setStyle("cart-panel");
        $this->container->append($panelLeft);

        $cartHeader = new DOM();
        $cartHeader->setText("Cart");
        $panelLeft->append($cartHeader);

        $cartSum = new DOM();
        $sum = 0;
        $panelLeft->append($cartSum);


        $panelRight = new DOM();
        $panelRight->setStyle("cart-filler");
        $this->container->append($panelRight);

        foreach ($cart as $itemCart) {
            //$item->toJson();
            $itemRow = new ItemCartRow($itemCart);
            $item = GetItem($itemCart->getID());
            $sum += $itemCart->getCount() * $item->getPrice();
            $panelRight->append($itemRow);
        }

        $cartSum->setText("Total: " . $sum);

        $buyButton = new Button();
        $buyButton->setText("Buy");
        $buyButton->setAttribute("margin-top", "auto")->setStyle("buy-button");
        $buyButton->setAction(function () use ($user) {
            BuyCart($user->getId());
        });
        $panelLeft->append($buyButton);

        return $this;
    }
}

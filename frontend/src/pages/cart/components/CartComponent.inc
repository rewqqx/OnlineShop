<?php

class CartComponent extends DOM
{
    private CartHeader $header;
    private Container $container;

    protected function createElement()
    {
        $this->header = new CartHeader();
        $this->container = new Container();
        $this->append($this->header)->append($this->container);

        $url = (empty($_SERVER['HTTPS']) ? 'http' : 'https') . "://$_SERVER[HTTP_HOST]$_SERVER[REQUEST_URI]";
        $parts = parse_url($url);
        parse_str($parts['query'], $query);

        $token = getTokenFromCookie();

        if (!$token->isValid()) {
            header('Location: /OnlineShop/frontend/src/index.php');
        }

        $this->container->setTag("form");
        $this->container->setStyle("profile-page");
        $this->container->setMethod("POST");

        $user = GetUser($token);

        $cart = GetCart($user->getId());

        $panel = new DOM();
        $panel->setStyle("cart-panel");
        $this->container->append($panel);

        $cartHeader = new DOM();
        $cartHeader->setStyle("cart-header")->setText("Cart");
        $panel->append($cartHeader);

        $horizontalContainer = new DOM();
        $horizontalContainer->setStyle("container-horizontal");
        $panel->append($horizontalContainer);

        $verticalItemLeft = new DOM();
        $verticalItemLeft->setStyle("container-vertical");
        $horizontalContainer->append($verticalItemLeft);

        $items = $cart->getItems();

        $itemsListHeader = new DOM();
        $itemsListHeader->setText("Item    Quantity   Price")->setStyle("cart");
        $verticalItemLeft->append($itemsListHeader);

//        print_r(GetItem($items[0]['item_id']));
        foreach ($items as $key => $item) {
            $itemsList = new DOM();
//            $item_name = GetItem($item['item_id'])->getName();
            $itemsList->setText($item['item_id']."    ".$item['quantity']."   ".$item['price'].'$')->setStyle("cart-item");
            $verticalItemLeft->append($itemsList);
        }



    }
}
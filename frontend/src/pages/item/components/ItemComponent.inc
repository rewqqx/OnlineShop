<?php

class ItemComponent extends DOM
{
    private ItemHeader $header;
    private Container $container;

    protected function createElement(): DOM
    {
        $this->header = new ItemHeader();
        $this->container = new Container();
        $this->append($this->header)->append($this->container);

        $url = (empty($_SERVER['HTTPS']) ? 'http' : 'https') . "://$_SERVER[HTTP_HOST]$_SERVER[REQUEST_URI]";
        $parts = parse_url($url);
        parse_str($parts['query'], $query);

        $item = GetItem($query['item_id']);

        $this->container->setTag("form");
        $this->container->setStyle("item-page");
        $this->container->setMethod("POST");

        $panel = new DOM();
        $panel->setStyle("item-panel");
        $this->container->append($panel);


        $itemHeader = new DOM();
        $itemHeader->setStyle("item-header")->setText($item->getName());
        $panel->append($itemHeader);

        $horizontalContainer = new DOM();
        $horizontalContainer->setStyle("container-horizontal");
        $panel->append($horizontalContainer);

        $verticalItemLeft = new DOM();
        $verticalItemLeft->setStyle("container-vertical");
        $horizontalContainer->append($verticalItemLeft);

        $icon = new DOM();
        $icon->setStyle("item-image-big");
        $verticalItemLeft->append($icon);

        $price = new DOM();
        $price->setText("Price: " . $item->getPrice() . "$")->setStyle("item-price");
        $verticalItemLeft->append($price);

        $verticalItemRight = new DOM();
        $verticalItemRight->setStyle("container-vertical");
        $horizontalContainer->append($verticalItemRight);


        $token = getTokenFromCookie();

        $confirm = new Button("ADD TO CART");
        $confirm->setAttribute("align-self", "center")->setAttribute("background", "#fafafa")->setAttribute("margin-top", "auto")->setAttribute("margin-bottom", "16px");
        $confirm->setAction(function () use ($item, $token) {
            $data = array("item_id" => $item->getID(), "user_id" => $token->getID(), "count" => 1);
            $cartItem = new CartItem($data);
            PutCartItem($cartItem);
        });

        $panel->append($confirm);
        return $this;
    }
}
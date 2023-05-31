<?php

class CartHeader extends DOM
{

    private Link $profile;
    private Text $title;

    protected function createElement(): DOM
    {
        $this->title = new Text("Cart");
        $this->title->setStyle("header-text");

        $buttonName = "Cart";
        $link = "/frontend/src/index.php";

        $this->profile = new Link($buttonName);
        $this->profile->setAttribute("margin-left", "auto");
        $this->profile->setAttribute("margin-right", "20px");

        $this->profile->setLink($link);

        $this->setStyle("header");

        $this->append($this->title);
        $this->append($this->profile);


        return $this;
    }
}

?>
<?php

class ItemComponent extends DOM
{
    private ItemHeader $header;
    private Container $container;

    protected function createElement()
    {
        $this->header = new ItemHeader();
        $this->container = new Container();
        $this->append($this->header)->append($this->container);


        $this->container->setTag("form");
        $this->container->setStyle("profile-page");
        $this->container->setMethod("POST");
    }
}
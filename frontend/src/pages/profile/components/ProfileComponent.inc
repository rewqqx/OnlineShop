<?php

class ProfileComponent extends DOM
{
    private ProfileHeader $header;
    private Container $container;

    protected function createElement()
    {
        $this->header = new ProfileHeader();
        $this->container = new Container();
        $this->append($this->header)->append($this->container);


        $this->container->setTag("form");
        $this->container->setStyle("profile-page");
        $this->container->setMethod("POST");
    }
}
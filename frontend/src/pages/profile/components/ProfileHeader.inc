<?php

class ProfileHeader extends DOM
{

    private Link $profile;
    private Text $title;

    protected function createElement()
    {
        $this->title = new Text("Profile");
        $this->title->setStyle("header-text");

        $buttonName = "Shop";
        $link = "/OnlineShop/frontend/src/index.php";

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
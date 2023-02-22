<?php

class Header extends DOM
{

    private Link $profile;
    private Text $title;

    protected function createElement()
    {
        $this->title = new Text("Online shop");
        $this->title->setStyle("header-text");

        $this->profile = new Link("Profile");
        $this->profile->setAttribute("margin-left", "auto");
        $this->profile->setAttribute("margin-right", "20px");
        $this->profile->setLink("./pages/auth/AuthPage.inc");

        $this->setStyle("header");

        $this->append($this->title);
        $this->append($this->profile);


        return $this;
    }
}

?>
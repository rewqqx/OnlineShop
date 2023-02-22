<?php

class Header extends DOM
{

    private Button $profile;
    private Text $title;

    protected function createElement()
    {
        $this->title = new Text("Online shop");
        $this->title->setStyle("header-text");

        $this->profile = new Button("Profile");
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
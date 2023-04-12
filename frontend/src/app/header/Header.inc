<?php

class Header extends DOM
{

    private Link $profile;
    private Text $title;

    protected function createElement()
    {
        $this->title = new Text("Online shop");
        $this->title->setStyle("header-text");

        $token = getTokenFromCookie();

        $buttonName = "Sign In";

        if ($token->isValid()) {
            $buttonName = "Profile";
        }

        $this->profile = new Link($buttonName);
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
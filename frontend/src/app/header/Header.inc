<?php

class Header extends DOM
{
    protected function createElement()
    {
        $this->setStyle("header");
        $this->setText("Online Shop");
        return $this;
    }
}

?>
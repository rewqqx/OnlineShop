<?php

class Toolbar extends DOM
{
    protected function createElement(): DOM
    {
        $this->setStyle("toolbar");
        return $this;
    }
}

?>
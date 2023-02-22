<?php

include "../toolbar/Toolbar.inc";
class Container extends DOM
{
    protected Toolbar $toolbar;

    protected function createElement()
    {
        $this->toolbar = new Toolbar();

        $this->setStyle("container");
        $this->append($this->toolbar);

        return $this;
    }
}

?>
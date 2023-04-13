<?php


class Container extends DOM
{
    protected Toolbar $toolbar;
    protected ItemGrid $grid;

    protected function createElement()
    {
        $this->setStyle("container");
        return $this;
    }
}

?>
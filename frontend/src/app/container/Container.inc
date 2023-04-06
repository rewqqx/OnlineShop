<?php

include "./app/toolbar/Toolbar.inc";
include "./app/grid/ItemGrid.inc";

class Container extends DOM
{
    protected Toolbar $toolbar;
    protected ItemGrid $grid;

    protected function createElement()
    {
        $this->setStyle("container");

        $this->toolbar = new Toolbar();
        $this->append($this->toolbar);

        $this->grid = new ItemGrid();
        $this->append($this->grid);

        return $this;
    }
}

?>
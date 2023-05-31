<?php

class ItemCartRow extends DOM
{

    private $item = null;
    private $itemCart = null;

    public function __construct($itemCart)
    {
        $this->itemCart = $itemCart;
        $this->item = GetItem($itemCart->getID());
        $this->setTag("div");
        $this->createElement();
    }

    protected function createElement(): DOM
    {
        parent::createElement();

        $this->setStyle("item-row");

        $desc = new DOM();
        $desc->setStyle("item-desc")->setText($this->item->getName() . ": " . $this->itemCart->getCount() . " x " . $this->item->getPrice());
        $this->append($desc);

        return $this;
    }


}

?>
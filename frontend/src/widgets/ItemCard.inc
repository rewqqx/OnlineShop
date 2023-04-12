<?php

class ItemCard extends DOM
{

    private Item|null $item = null;

    public function __construct($item)
    {
        $this->item = $item;
        $this->setTag("div");
        $this->createElement();
    }

    protected function createElement()
    {
        parent::createElement();

        $this->setStyle("item-card");

        $image = new DOM();
        $image->setStyle("item-image");
        $this->append($image);

        $desc = new DOM();
        $desc->setStyle("item-desc")->setText($this->item->getName() . " - " . $this->item->getPrice() . " $");
        $this->append($desc);


        return $this;
    }


}

?>
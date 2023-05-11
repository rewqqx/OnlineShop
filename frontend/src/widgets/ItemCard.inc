<?php

class ItemCard extends DOM
{

    private $item = null;

    public function __construct($item)
    {
        $this->item = $item;
        $this->setTag("a");
        $this->createElement();
    }

    protected function createElement(): DOM
    {
        parent::createElement();

        $this->setStyle("item-card");
        $image = new Image("http://" . getURL("IMAGE_HOST") . ":9081/image/" . $this->item->getImageIDs()[0]);
        $image->setStyle("item-image");
        $this->append($image);

        $desc = new DOM();
        $desc->setStyle("item-desc")->setText($this->item->getName() . " - " . $this->item->getPrice() . " $");
        $this->append($desc);


        return $this;
    }


}

?>
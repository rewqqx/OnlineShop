<?php

class ItemGrid extends DOM
{
    protected function createElement(): DOM
    {
        $this->setStyle("grid");


        foreach (GetItems() as $item) {
            $link = "/frontend/src/pages/item/ItemPage.php?item_id=" . $item->getId();
            $item->toJson();
            $itemCard = new ItemCard($item);
            $itemCard->setLink($link);
            $this->append($itemCard);
        }

        return $this;
    }
}

?>
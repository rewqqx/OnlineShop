<?php

class ItemGrid extends DOM
{
    protected function createElement(): DOM
    {
        $this->setStyle("grid");


        foreach (GetItems() as $item) {
            $link = "/OnlineShop/frontend/src/pages/item/ItemPage.inc?item_id=" . $item->getId();
            $item->toJson();
            $itemCard = new ItemCard($item);
            $itemCard->setLink($link);
            $this->append($itemCard);
        }

        return $this;
    }
}

?>
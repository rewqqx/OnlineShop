<?php

class ItemGrid extends DOM
{
    protected function createElement()
    {
        $this->setStyle("grid");


        foreach (GetItems() as $item) {
            $item->toJson();
            $itemCard = new ItemCard($item);
            $this->append($itemCard);
        }

        return $this;
    }
}

?>
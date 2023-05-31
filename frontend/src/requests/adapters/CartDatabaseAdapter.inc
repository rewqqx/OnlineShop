<?php

require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/objects/CartItem.inc");

function GetCartItems($id): array
{
    $response = GETRequest("http://" . getURL("SERVER_HOST") . ":9080/cart/" . $id);

    if ($response["status"] != "200") {
        return array();
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["cart"];

    $result = array();

    if ($response == null) {
        return $result;
    }

    foreach ($response as $value) {
        $result[] = new CartItem($value);
    }

    return $result;
}

function PutCartItem($item): bool
{
    $response = PUTRequest("http://" . getURL("SERVER_HOST") . ":9080/cart/", $item->toJson());

    if ($response["status"] != "200") {
        return false;
    }

    return true;
}
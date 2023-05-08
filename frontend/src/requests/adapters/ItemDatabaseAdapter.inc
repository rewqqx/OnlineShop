<?php

require_once( "./objects/Item.inc");


function GetItems(): array
{
    $response = GETRequest("http://" . getURL("SERVER_HOST") . ":8080/items/");

    if ($response["status"] != "200") {
        return array();
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["items"];

    $result = array();

    foreach ($response as $value) {
        $result[] = new Item($value);
    }

    return $result;
}

function GetItem($id): Item
{
    $response = GETRequest("http://" . getURL("SERVER_HOST") . ":8080/items/" . $id);

    if ($response["status"] != "200") {
        return array();
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["item"];

    return new Item($response);
}

?>
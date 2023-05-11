<?php

require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/objects/Item.inc");


function GetItems(): array
{

    $url = (empty($_SERVER['HTTPS']) ? 'http' : 'https') . "://$_SERVER[HTTP_HOST]$_SERVER[REQUEST_URI]";
    $parts = parse_url($url);
    parse_str($parts['query'], $query);

    if (isset($query['tag_id'])) {
        $selectedID = $query['tag_id'];
        $response = POSTRequest("http://" . getURL("SERVER_HOST") . ":9080/items/", "{\"tag\": " . $selectedID . "}");
    } else {
        $response = GETRequest("http://" . getURL("SERVER_HOST") . ":9080/items/");
    }


    if ($response["status"] != "200") {
        return array();
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["items"];

    $result = array();

    if ($response === null) {
        return $result;
    }

    foreach ($response as $value) {
        $result[] = new Item($value);
    }

    return $result;
}

function GetItem($id): Item
{
    $response = GETRequest("http://" . getURL("SERVER_HOST") . ":9080/items/" . $id);

    if ($response["status"] != "200") {
        return array();
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["item"];

    return new Item($response);
}

?>
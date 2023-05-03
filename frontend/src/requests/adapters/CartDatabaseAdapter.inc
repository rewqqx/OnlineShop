<?php
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/objects/Cart.inc");
function GetCart($user_id): Cart
{
    $response = GETRequest("http://127.0.0.1:8080/cart/".$user_id);

    if ($response["status"] != "200") {
        return array();
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["cart"];

    return new Cart($response);

}
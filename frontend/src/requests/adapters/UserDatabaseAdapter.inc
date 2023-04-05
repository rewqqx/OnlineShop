<?php

include "./objects/Token.inc";
function GetUserToken($auth): ?Token
{
    $response = POSTRequest("http://127.0.0.1:8080/auth", $auth->toJSON());

    if ($response["status"] != "200") {
        return null;
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["token"];

    return new Token($response);
}

?>
<?php

include "./objects/Token.inc";
include "./objects/User.inc";
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

function GetUser($token): ?User
{
    $headers = HeadersJSON();
    $headers[] = "token: " . $token->getToken();
    $response = RequestSpecifyHeaders("http://127.0.0.1:8080/users/" . $token->getID(), "{}", $headers, "POST");

    if ($response["status"] != "200") {
        return null;
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["user"];

    return new User($response);
}

?>
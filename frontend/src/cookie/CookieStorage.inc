<?php

require_once("./objects/Token.inc");

$KEY = "shop_token";
$DAY = 60 * 60 * 24;
function getTokenFromCookie(): Token
{
    global $KEY;

    if (!isset($_COOKIE[$KEY])) {
        return new Token(array());
    }

    $tokenJSON = $_COOKIE[$KEY];
    return new Token(json_decode($tokenJSON, true));
}

function setTokenCookie(Token $token)
{
    global $DAY, $KEY;
    setcookie($KEY, $token->toJson(), time() + $DAY, '/');
}
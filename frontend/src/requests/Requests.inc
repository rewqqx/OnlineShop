<?php

function HeadersJSON(): array
{
    $headers = array();

    $headers[] = "Content-Type: application/x-www-form-urlencoded";
    $headers[] = "Content-Type: application/json";
    $headers[] = "Accept: application/json";
    $headers[] = "Access-Control-Allow-Headers: Content-Type";

    return $headers;
}

function GetURL($key)
{
    $ip = getenv($key, true) ?: "127.0.0.1";
    return $ip;
}

function GETRequest($url)
{
    return RequestSpecifyHeaders($url, "", HeadersJSON(), "GET");
}

function PUTRequest($url, $body)
{
    return RequestSpecifyHeaders($url, $body, HeadersJSON(), "PUT");
}

function POSTRequest($url, $body)
{
    return RequestSpecifyHeaders($url, $body, HeadersJSON(), "POST");
}

function RequestSpecifyHeaders($url, $body, $headers, $type)
{
    $options = array('http' => array(
        'method' => $type,
        'header' => $headers,
        'ignore_errors' => true,
        'content' => $body
    ));
    $context = stream_context_create($options);
    $response = file_get_contents($url, false, $context);

    $status_line = $http_response_header[0];

    preg_match('{HTTP\/\S*\s(\d{3})}', $status_line, $match);

    $status = $match[1];

    $output["status"] = $status;
    $output["response"] = $response;

    return $output;
}


?>
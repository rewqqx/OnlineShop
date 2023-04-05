<?php

class Request
{

    public static function HeadersJSON(): array
    {
        $headers = array();

        $headers[] = "Content-Type: application/x-www-form-urlencoded";
        $headers[] = "Content-Type: application/json";
        $headers[] = "Accept: application/json";
        $headers[] = "Access-Control-Allow-Headers: Content-Type";

        return $headers;
    }

    public static function GETRequest($url)
    {
        return self::GETRequestSpecifyHeaders($url, self::HeadersJSON());
    }

    public static function POSTRequest($url, $body)
    {
        return self::POSTRequestSpecifyHeaders($url, $body, self::HeadersJSON());
    }

    public static function GETRequestSpecifyHeaders($url, $headers)
    {
        $q = array("q" => "{}");
        $options = array('http' => array(
            'method' => "GET",
            'header' => $headers,
            'content' => http_build_query($q)
        ));
        $context = stream_context_create($options);
        $response = file_get_contents($url, false, $context);
        return $response;
    }

    public static function POSTRequestSpecifyHeaders($url, $body, $headers)
    {
        $options = array('http' => array(
            'method' => "POST",
            'header' => $headers,
            'content' => $body
        ));
        $context = stream_context_create($options);
        $response = file_get_contents($url, false, $context);
        return $response;
    }
}

?>
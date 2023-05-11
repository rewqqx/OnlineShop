<?php

require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/objects/Tag.inc");

function GetTags(): array
{
    $response = GETRequest("http://" . getURL("SERVER_HOST") . ":9080/tags/");

    if ($response["status"] != "200") {
        return array();
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["tags"];

    $result = array();

    foreach ($response as $value) {
        $result[] = new Tag($value);
    }

    return $result;
}

function GetTag($id): Tag
{
    $response = GETRequest("http://" . getURL("SERVER_HOST") . ":9080/tags/" . $id);

    if ($response["status"] != "200") {
        return array();
    }

    $response = $response["response"];

    $response = json_decode($response, true);

    $response = $response["tag"];

    return new Tag($response);
}
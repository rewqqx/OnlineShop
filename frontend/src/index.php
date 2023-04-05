<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>

<?php
include "./app/App.inc";

$url = "http://127.0.0.1:8080/items/";
$q   = array("q"=>"PHP HTTP request");
$page = file_get_contents($url . '?' . http_build_query($q));

echo $page;

$app = new App();
echo $app->getDOM();

?>
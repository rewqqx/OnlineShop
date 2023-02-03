<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>

<?php

include "./elements/DOM.inc";
/*
$dbHost = 'localhost';
$dbPort = 49153;
$user = 'postgres';
$pass = 'postgrespw';
$dbConn = pg_connect("host=$dbHost port=$dbPort dbname=postgres user=$user password=$pass")
or die('Не удалось соединиться: ' . pg_last_error());

$sql = "SELECT * FROM test.test_1";

$pg_query = pg_query($dbConn, $sql);

$arr = pg_fetch_all($pg_query);
print_r($arr);*/

$header = new DOM();
$header->setStyle("header")->setText("Online Shop");

$toolbar = new DOM();
$toolbar->setStyle("toolbar");

$container = new DOM();
$container->setStyle("container")->append($toolbar);


echo $header->getDOM();
echo $container->getDOM();

?>
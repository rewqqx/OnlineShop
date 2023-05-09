<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>
<link href='styles/Item.css' rel='stylesheet' type='text/css'>
<link href='styles/Grid.css' rel='stylesheet' type='text/css'>

<?php

ini_set('display_errors', '1');
ini_set('display_startup_errors', '1');
error_reporting(E_ALL);

require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/app/App.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/requests/Requests.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/objects/Auth.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/widgets/ItemCard.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/requests/adapters/UserDatabaseAdapter.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/requests/adapters/ItemDatabaseAdapter.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/app/toolbar/Toolbar.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/app/grid/ItemGrid.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/cookie/CookieStorage.inc");

//echo GetUser($token)->toJson();

//echo $itemCard->getDOM();

$app = new App();
echo $app->getDOM();

?>
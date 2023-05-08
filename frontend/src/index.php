<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>
<link href='styles/Item.css' rel='stylesheet' type='text/css'>
<link href='styles/Grid.css' rel='stylesheet' type='text/css'>

<?php


ini_set('display_errors', '1');
ini_set('display_startup_errors', '1');
error_reporting(E_ALL);

echo ".index.php";

require_once("./app/App.inc");
require_once("./requests/Requests.inc");
require_once("./objects/Auth.inc");
require_once("./widgets/ItemCard.inc");
require_once("./requests/adapters/UserDatabaseAdapter.inc");
require_once("./requests/adapters/ItemDatabaseAdapter.inc");
require_once("./app/toolbar/Toolbar.inc");
require_once("./app/grid/ItemGrid.inc");
require_once("./cookie/CookieStorage.inc");

//echo GetUser($token)->toJson();

//echo $itemCard->getDOM();

$app = new App();
echo $app->getDOM();

?>
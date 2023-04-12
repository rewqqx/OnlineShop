<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>
<link href='styles/Item.css' rel='stylesheet' type='text/css'>
<link href='styles/Grid.css' rel='stylesheet' type='text/css'>

<?php
include "./app/App.inc";
include "./requests/Requests.inc";
include "./objects/Auth.inc";
include "./widgets/ItemCard.inc";
include "./requests/adapters/UserDatabaseAdapter.inc";
include "./requests/adapters/ItemDatabaseAdapter.inc";

//echo GetUser($token)->toJson();

//echo $itemCard->getDOM();

$app = new App();
echo $app->getDOM();

?>
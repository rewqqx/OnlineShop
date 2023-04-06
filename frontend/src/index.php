<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>

<?php
include "./app/App.inc";
include "./requests/Requests.inc";
include "./objects/Auth.inc";
include "./requests/adapters/UserDatabaseAdapter.inc";
include "./requests/adapters/ItemDatabaseAdapter.inc";

$auth = new Auth("admin@mail.ru", "admin");
$token = GetUserToken($auth);
echo GetUser($token)->toJson();
/*
foreach (GetItems() as $item) {
    echo $item->toJson();
}*/

echo GetItem(2)->toJson();


$app = new App();
echo $app->getDOM();

?>
<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>

<?php
include "./app/App.inc";
include "./requests/Requests.inc";
include "./objects/Auth.inc";
include "./requests/adapters/UserDatabaseAdapter.inc";

$auth = new Auth("admin@mail.ru", "admin");
$token = GetUserToken($auth);
echo GetUser($token)->toJson();
//echo Request::GETRequest("http://127.0.0.1:8080/items/");


$app = new App();
echo $app->getDOM();

?>
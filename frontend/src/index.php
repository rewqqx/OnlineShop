<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>

<?php
include "./app/App.inc";
include "./requests/Requests.inc";
include "./objects/Auth.inc";

$auth = new \user\Auth("admin@mail.ru", "admin");
echo $auth->toJson();
echo Request::POSTRequest("http://127.0.0.1:8080/auth", $auth->toJson());
//echo Request::GETRequest("http://127.0.0.1:8080/items/");


$app = new App();
echo $app->getDOM();

?>
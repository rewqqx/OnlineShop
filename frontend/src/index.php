<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>

<?php
include "./app/App.inc";
include "./requests/Requests.inc";

echo Request::POSTRequest("http://127.0.0.1:8080/auth/", "{
    \"mail\": \"admin@mail.ru\",
    \"password\": \"admin\"}");
echo Request::GETRequest("http://127.0.0.1:8080/items/");


$app = new App();
echo $app->getDOM();

?>
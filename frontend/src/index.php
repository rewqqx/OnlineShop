<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>

<?php

include "./app/App.inc";

$app = new App();
echo $app->getDOM();

?>
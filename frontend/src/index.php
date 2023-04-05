<link href='styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='styles/Default.css' rel='stylesheet' type='text/css'>
<link href='styles/Auth.css' rel='stylesheet' type='text/css'>

<?php

include "./database/Database.inc";
include "./app/App.inc";

\database\Database::getDatabase();

$app = new App();
echo $app->getDOM();

?>
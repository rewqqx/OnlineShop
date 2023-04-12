<link href='../../styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='../../styles/Default.css' rel='stylesheet' type='text/css'>
<link href='../../styles/Auth.css' rel='stylesheet' type='text/css'>

<?php
include $_SERVER['DOCUMENT_ROOT'] . "\\frontend\src" ."\\requests\Requests.inc";
include $_SERVER['DOCUMENT_ROOT'] . "\\frontend\src" . "\\objects\Auth.inc";
include $_SERVER['DOCUMENT_ROOT'] . "\\frontend\src" . "\\elements\DOM.inc";
include $_SERVER['DOCUMENT_ROOT'] . "\\frontend\src" . "\\requests\adapters\UserDatabaseAdapter.inc";

include "./components/AuthComponent.inc";

$auth = new AuthComponent();
echo $auth->getDOM();
?>
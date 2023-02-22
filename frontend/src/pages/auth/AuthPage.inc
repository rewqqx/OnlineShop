<link href='../../styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='../../styles/Default.css' rel='stylesheet' type='text/css'>
<link href='../../styles/Auth.css' rel='stylesheet' type='text/css'>

<?php
include "../../elements/DOM.inc";
include "./components/Auth.inc";

$auth = new Auth();
echo $auth->getDOM();
?>
<link href='../../styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='../../styles/Default.css' rel='stylesheet' type='text/css'>
<link href='../../styles/Auth.css' rel='stylesheet' type='text/css'>

<?php
require_once( "./requests/Requests.inc");
require_once( "./objects/Auth.inc");
require_once( "./elements/DOM.inc");
require_once( "./cookie/CookieStorage.inc");
require_once( "./requests/adapters/UserDatabaseAdapter.inc");

require_once("./components/AuthComponent.inc");

$auth = new AuthComponent();
echo $auth->getDOM();
?>
<link href='../../styles/DOM.css' rel='stylesheet' type='text/css'>
<link href='../../styles/Default.css' rel='stylesheet' type='text/css'>
<link href='../../styles/Cart.css' rel='stylesheet' type='text/css'>
<link href='../../styles/Item.css' rel='stylesheet' type='text/css'>

<?php
ini_set('display_errors', '1');
ini_set('display_startup_errors', '1');
error_reporting(E_ALL);

require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/requests/Requests.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/objects/Auth.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/elements/DOM.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/cookie/CookieStorage.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/requests/adapters/UserDatabaseAdapter.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/requests/adapters/CartDatabaseAdapter.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/requests/adapters/ItemDatabaseAdapter.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/widgets/ItemCartRow.inc");

require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/elements/DOM.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/app/container/Container.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/pages/cart/components/CartComponent.inc");
require_once($_SERVER['DOCUMENT_ROOT'] . "/frontend/src" . "/pages/cart/components/CartHeader.inc");

$cart = new CartComponent();
echo $cart->getDOM();
?>
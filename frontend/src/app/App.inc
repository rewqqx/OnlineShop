<?php
include "./elements/DOM.inc";
include "./app/container/Container.inc";
include "./app/header/Header.inc";

class App extends DOM
{

    private Header $header;
    private Container $container;

    protected function createElement()
    {
        $this->header = new Header();
        $this->container = new Container();

        $this->append($this->header)->append($this->container);

        return $this;
    }

}

?>
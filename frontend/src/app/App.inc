<?php
include "./container/Container.inc";
include "./header/Header.inc";

class App extends DOM
{

    private Header $header;
    private Container $container;

    protected function createElement()
    {
        $this->header = new Header();
        $this->container = new Container();

        $this->append($this->header)->append($this . $this->container);

        return $this;
    }

}

?>
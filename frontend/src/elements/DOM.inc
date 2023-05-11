<?php


class DOM
{

    protected ?string $uid = null;
    protected string $tag = "div";
    protected ?string $id = null;
    protected ?string $href = null;
    protected ?string $url = null;
    protected ?string $type = null;
    protected ?string $style = null;
    protected ?string $value = null;

    protected ?string $method = null;

    protected ?string $name = null;
    protected ?string $text = null;
    protected array $attributes = [];
    protected array $children = [];

    function __construct()
    {
        $this->setTag("div");
        $this->createElement();
    }


    protected function createElement(): DOM
    {
        return $this;
    }

    protected function setTag(string $tag): DOM
    {
        $this->uid = uniqid($this->tag . "_");
        $this->tag = $tag;
        return $this;
    }

    protected function setValue(string $value): DOM
    {
        $this->value = $value;
        return $this;
    }

    protected function setMethod(string $method): DOM
    {
        $this->method = $method;
        return $this;
    }

    protected function setName(string $name): DOM
    {
        $this->name = $this->tag . "_" . strtolower(strtr($name, array(' ' => '_')));
        return $this;
    }

    public function setStyle(string $style): DOM
    {
        $this->style = $style;
        return $this;
    }

    public function setText(string $text): DOM
    {
        $this->text = $text;
        return $this;
    }


    public function setID(string $id): DOM
    {
        $this->id = $id;
        return $this;
    }

    public function setAttribute(string $name, string $value): DOM
    {
        $this->attributes[$name] = $value;
        return $this;
    }


    public function append(DOM $child): DOM
    {
        array_push($this->children, $child);
        return $this;
    }

    public function setLink(string $link): DOM
    {
        $this->href = $link;
        return $this;
    }

    public function setURL(string $link): DOM
    {
        $this->url = $link;
        return $this;
    }

    public function setType(string $type): DOM
    {
        $this->type = $type;
        return $this;
    }

    public function getDOM(): string
    {
        $result = "<" . $this->tag;

        if (!is_null($this->id)) {
            $result = $result . ' id = "' . $this->id . '"';
        }

        if (!is_null($this->name)) {
            $result = $result . ' name = "' . $this->name . '"';
        }

        if (!is_null($this->method)) {
            $result = $result . ' method = "' . $this->method . '"';
        }

        if (!is_null($this->value)) {
            $result = $result . ' value = "' . $this->value . '"';
        }

        if (!is_null($this->href)) {
            $result = $result . ' href = "' . $this->href . '"';
        }

        if (!is_null($this->url)) {
            $result = $result . ' src = "' . $this->url . '"';
        }

        if (!is_null($this->type)) {
            $result = $result . ' type = "' . $this->type . '"';
        }

        if (!is_null($this->style)) {
            $result = $result . ' class = "' . $this->style . '"';
        }

        if (count($this->attributes) > 0) {
            $result = $result . ' style="';
            foreach ($this->attributes as $key => $value) {
                $result = $result . $key . ":" . $value . ";";
            }
            $result = $result . '"';
        }

        $result = $result . ">";

        if (!is_null($this->text)) {
            $result = $result . $this->text;
        }

        foreach ($this->children as $key => $value) {
            $result = $result . "\n" . $value->getDOM();
        }

        $result = $result . "</" . $this->tag . ">\n";

        return $result;
    }

}


class Text extends DOM
{
    function __construct($text = "")
    {
        $this->setTag("div");
        $this->text = $text;
        $this->createElement();
    }
}


class Link extends DOM
{
    public Text $textElement;

    function __construct($text = "")
    {
        $this->setTag("a");

        $this->textElement = new Text($text);
        $this->textElement->setStyle("button-text");

        $this->setStyle("button");
        $this->append($this->textElement);
    }
}

class Image extends DOM
{

    function __construct($url = "")
    {
        $this->setTag("img");
        $this->setURL($url);
    }
}

class Button extends DOM
{

    // TODO: Попробуй пришить тут лямбду функцию.
    // Если не получится можно сделать объект Command и метод Execute у него переопределять и подавать его на вход
    public ?Closure $action = null;

    function __construct($text = "")
    {
        $this->setTag("button");
        $this->setValue($text);
        $this->setName($text);

        $this->createElement();
    }

    protected function createElement(): DOM
    {
        $this->setType("submit");
        $this->setStyle("button");

        return $this;
    }

    public function setAction($action): DOM
    {
        $this->action = $action;
        $this->processClickEvent();

        return $this;
    }

    protected function processClickEvent(): void
    {
        if (array_key_exists($this->name, $_POST)) {
            $this->clickEvent();
            return;
        }

        $url = (empty($_SERVER['HTTPS']) ? 'http' : 'https') . "://$_SERVER[HTTP_HOST]$_SERVER[REQUEST_URI]";

        if ($url == null) {
            return;
        }

        $parts = parse_url($url);

        if (!array_key_exists("query", $parts)) {
            return;
        }

        parse_str($parts['query'], $query);

        if (array_key_exists($this->name, $query)) {
            $this->clickEvent();
        }
    }

    private function clickEvent(): void
    {
        if ($this->action == null) {
            echo 'Name [' . $this->name . "] is clicked";
        } else {
            ($this->action)();
        }
    }
}

class Input extends DOM
{
    function __construct($text = "")
    {
        $this->setTag("input");
        $this->setName($text);
        $this->setStyle("input");

        $this->createElement();
    }

}

class Icon extends DOM
{

    protected ?string $path = null;

    function __construct($path = "")
    {
        $this->setTag("img");
        $this->setURL("/frontend/resources/icons/" . $path);
        $this->createElement();
    }


}

class IconButton extends Button
{
    protected ?string $path = null;

    function __construct($path = "")
    {
        $this->path = $path;
        parent::__construct();
    }

    protected function createElement(): DOM
    {
        parent::createElement();
        $this->setAttribute("background", "transparent");
        $icon = new Icon($this->path);
        $icon->setStyle("white-filter");
        $this->append($icon);

        return $this;
    }
}

?>
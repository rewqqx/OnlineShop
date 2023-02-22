<?php


class DOM
{

    protected ?string $uid = null;
    protected string $tag = "div";
    protected ?string $id = null;
    protected ?string $href = null;
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


    protected function createElement()
    {
        return $this;
    }

    protected function setTag(string $tag)
    {
        $this->uid = uniqid($this->tag . "_");
        $this->tag = $tag;
    }

    protected function setValue(string $value)
    {
        $this->value = $value;
        return $this;
    }

    protected function setMethod(string $method)
    {
        $this->method = $method;
        return $this;
    }

    protected function setName(string $name)
    {
        $this->name = $this->tag . "_" . strtolower(strtr($name, array(' ' => '_')));
        return $this;
    }

    public function setStyle(string $style)
    {
        $this->style = $style;
        return $this;
    }

    public function setText(string $text)
    {
        $this->text = $text;
        return $this;
    }


    public function setID(string $id)
    {
        $this->id = $id;
        return $this;
    }

    public function setAttribute(string $name, string $value)
    {
        $this->attributes[$name] = $value;
        return $this;
    }

    public function append(DOM $child)
    {
        array_push($this->children, $child);
        return $this;
    }

    public function setLink(string $link)
    {
        $this->href = $link;
        return $this;
    }

    public function setType(string $type)
    {
        $this->type = $type;
        return $this;
    }

    public function getDOM()
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

class Button extends DOM
{

    // TODO: Попробуй пришить тут лямбду функцию.
    // Если не получится можно сделать объект Command и метод Execute у него переопределять и подавать его на вход
    public ?Closure $action = null;

    function __construct($text = "")
    {
        $this->setTag("input");
        $this->setValue($text);
        $this->setName($text);

        $this->createElement();
    }

    protected function createElement()
    {
        $this->setType("submit");
        $this->setStyle("button");
        $this->processClickEvent();

        return $this;
    }

    protected function processClickEvent()
    {
        if (array_key_exists($this->name, $_POST)) {
            $this->clickEvent();
        }
    }

    private function clickEvent()
    {
        echo 'action: ' . $this->action;
        if ($this->action == null) {
            echo 'Name [' . $this->name . "] is clicked";
        } else {
            $this->action();
        }
    }
}

class Input extends DOM
{
    function __construct()
    {
        $this->setTag("input");
        $this->setStyle("input");
    }
}

?>
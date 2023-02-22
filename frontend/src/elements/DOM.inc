<?php


class DOM
{

    protected string $tag = "div";
    protected ?string $id = null;
    protected ?string $href = null;
    protected ?string $type = null;
    protected ?string $style = null;
    protected ?string $text = null;
    protected array $attributes = [];
    protected array $children = [];

    function __construct()
    {
        $this->createElement();
    }


    protected function createElement()
    {
        $this->setTag("div");
        return $this;
    }

    protected function setTag(string $tag)
    {
        $this->tag = $tag;
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
    function __construct($text)
    {
        $this->text = $text;
        $this->createElement();
    }
}


class Button extends DOM
{
    function __construct($text)
    {
        $textElement = new Text($text);
        $textElement->setStyle("button-text");

        $this->setTag("a");
        $this->setStyle("button");
        $this->append($textElement);
    }
}

?>
<?php


class DOM
{

    protected string $tag = "div";
    protected ?string $id = null;
    protected ?string $style = null;

    protected ?string $text = null;
    protected array $attributes = [];
    protected array $children = [];

    function __construct()
    {
        $this->setTag();
    }

    protected function setTag()
    {
        $this->tag = "div";
    }

    public function setStyle(string $style)
    {
        $this->$style = $style;
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

    public function getDOM()
    {
        $result = "<" . $this->tag;

        if (!is_null($this->id)) {
            $result = $result . ' id = "' . $this->id . '"';
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

?>
<?php

class Toolbar extends DOM
{
    protected function createElement(): DOM
    {
        $this->setStyle("toolbar");

        $tags = GetTags();

        $url = (empty($_SERVER['HTTPS']) ? 'http' : 'https') . "://$_SERVER[HTTP_HOST]$_SERVER[REQUEST_URI]";
        $parts = parse_url($url);

        if (isset($parts['query'])) {
            parse_str($parts['query'], $query);
        }

        if (isset($query['tag_id'])) {
            $selectedID = $query['tag_id'];
        } else {
            $selectedID = -1;
        }
        foreach ($tags as $tag) {
            $tagDOM = new Link();
            if ($selectedID != $tag->getId()) {
                $tagDOM->setStyle("tag-link");
            } else {
                $tagDOM->setStyle("tag-link-selected");
            }
            $tagDOM->setText($tag->getName())->setLink("?tag_id=" . $tag->getId());
            $this->append($tagDOM);
        }

        return $this;
    }
}

?>
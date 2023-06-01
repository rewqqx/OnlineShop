import * as item from "../../adapters/item/Item.js";

const currentUrl = window.location.href;
const parsedUrl = new URL(currentUrl);
const id = parsedUrl.searchParams.get('id');


const inputName = document.getElementById("name");
const inputPrice = document.getElementById("price");
const inputImages = document.getElementById("images");
const inputDescription = document.getElementById("description");
const inputTags = document.getElementById("tags");

if (id !== -1) {
    const data = item.getRowByID(id);

    inputName.value = data.name;
    inputPrice.value = data.price;
    inputTags.value = data.tag_ids;
    inputImages.value = data.image_ids;
    inputDescription.value = data.description;
}

const save = document.getElementById("save");
save.onclick = function () {
    const name = inputName.value;
    const price = inputPrice.value;
    const description = inputDescription.value;
    const images = inputImages.value;
    const tags = inputTags.value;

    if (id !== -1) {
        item.createRow({name: name, price: price, description: description, images_ids: images, tag_ids: tags})
    } else {

    }

    window.location.href = "./../itemPage/ListItemPage.html";
}
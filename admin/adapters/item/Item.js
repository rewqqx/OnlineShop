import * as library from "./../../library/RequestLibrary.js"
export class Item {
#description;
#images;
#tags;
#id;
#name;
#price;
constructor(json) {
this.#price = json.price;
this.#description = json.description;
this.#images = json.images;
this.#tags = json.tags;
this.#id = json.id;
this.#name = json.name;
}
getImages() {
return this.#images;
}
getTags() {
return this.#tags;
}
getID() {
return this.#id;
}
getName() {
return this.#name;
}
getPrice() {
return this.#price;
}
getDescription() {
return this.#description;
}
}
export function getRows() {
   const response = library.httpGet("http://127.0.0.1:9080/items/");
   const json = JSON.parse(response);
   return json.items;
}

export function getRowByID(id) {
   const response = library.httpGet("http://127.0.0.1:9080/items/" + id);
   const json = JSON.parse(response);
   return json.item;
}

export function updateRow(item) {
   //  Implement me! 
}

export function createRow(item) {
   //  Implement me! 
}
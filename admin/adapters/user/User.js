import * as library from "./../../library/RequestLibrary.js"
export class User {
#password;
#role;
#surname;
#patronymic;
#birthdate;
#mail;
#sex;
#id;
#name;
#phone;
constructor(json) {
this.#surname = json.surname;
this.#patronymic = json.patronymic;
this.#birthdate = json.birthdate;
this.#password = json.password;
this.#role = json.role;
this.#id = json.id;
this.#name = json.name;
this.#phone = json.phone;
this.#mail = json.mail;
this.#sex = json.sex;
}
getPassword() {
return this.#password;
}
getRole() {
return this.#role;
}
getSurname() {
return this.#surname;
}
getPatronymic() {
return this.#patronymic;
}
getBirthdate() {
return this.#birthdate;
}
getMail() {
return this.#mail;
}
getSex() {
return this.#sex;
}
getID() {
return this.#id;
}
getName() {
return this.#name;
}
getPhone() {
return this.#phone;
}
}
export function getRows() {
   const response = library.httpGet("http://127.0.0.1:9080/users/", "a4c3eed9907e10e89ca2df38af8d43d59a57c1d4cdd671a64d84afb2f66ee9ec");
   const json = JSON.parse(response);
   return json.users;
}

export function getRowByID(id) {
   const response = library.httpGet("http://127.0.0.1:9080/users/" + id, "a4c3eed9907e10e89ca2df38af8d43d59a57c1d4cdd671a64d84afb2f66ee9ec");
   const json = JSON.parse(response);
   return json.user;
}

export function updateRow(item) {
   //  Implement me! 
}

export function createRow(item) {
   //  Implement me! 
}
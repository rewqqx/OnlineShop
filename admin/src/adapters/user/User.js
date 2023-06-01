import * as library from "../../library/RequestLibrary.js"

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
    const response = library.httpGet("http://127.0.0.1:9080/users/", "680ee3efa31e13b750bcb34874b9e89390b8a5de5b633bc9e086a306cae54d33");
    const json = JSON.parse(response);
    return json.users;
}

export function getRowByID(id) {
    const response = library.httpGet("http://127.0.0.1:9080/users/" + id, "680ee3efa31e13b750bcb34874b9e89390b8a5de5b633bc9e086a306cae54d33");
    const json = JSON.parse(response);
    return json.user;
}

export function updateRow(item) {
    library.httpPost("http://127.0.0.1:9080/users/update/" + item.id, item);
}

export function createRow(item) {
    library.httpPost("http://127.0.0.1:9080/users/create", item);
}

export function deleteRow(id) {
    library.httpDelete("http://127.0.0.1:9080/users/delete/" + id);
}
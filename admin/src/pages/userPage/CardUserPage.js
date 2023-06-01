import * as user from "../../adapters/user/User.js";

const currentUrl = window.location.href;
const parsedUrl = new URL(currentUrl);
const id = parsedUrl.searchParams.get('id');


const inputName = document.getElementById("name");
const inputSurname = document.getElementById("surname");
const inputPatronymic = document.getElementById("patronymic");
const inputMail = document.getElementById("mail");
const inputPassword = document.getElementById("password");
const inputPhone = document.getElementById("phone");
const inputGender = document.getElementById("gender");
const inputRole = document.getElementById("role");
//const inputBirthday = document.getElementById("role");

if (id !== -1) {
    const data = user.getRowByID(id);

    inputName.value = data.user_name;
    inputSurname.value = data.user_surname;
    inputPatronymic.value = data.user_patronymic;
    inputMail.value = data.mail;
    inputPassword.value = data.password_hash;
    inputPhone.value = data.phone;
    inputGender.value = data.sex;
    inputRole.value = data.role_id;

}

const save = document.getElementById("save");
save.onclick = function () {
    const name = inputName.value;
    const surname = inputSurname.value;
    const patronymic = inputPatronymic.value;
    const phone = inputPhone.value;
    const mail = inputMail.value;
    const roleID = inputRole.value;
    const password = inputPassword.value;

    if (id !== -1) {
        user.createRow({
            id: -1,
            user_name: name,
            user_surname: surname,
            user_patronymic: patronymic,
            password_hash: password,
            phone: phone,
            birthdate: null,
            mail: mail,
            role_id: roleID,
            token: ""
        })
        window.location.href = "./../itemPage/ListItemPage.html";
    } else {
        user.updateRow(id, {
            user_name: name,
            user_surname: surname,
            user_patronymic: patronymic,
            phone: phone,
            birthdate: null,
            mail: mail,
            role_id: roleID,
        })
        window.location.href = currentUrl;
    }
}
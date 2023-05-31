export function httpGet(url, token) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", url, false); // false for synchronous request

    xmlHttp.setRequestHeader("token", token);

    xmlHttp.send();
    return xmlHttp.responseText;
}
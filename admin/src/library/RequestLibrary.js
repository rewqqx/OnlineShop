export function httpGet(url, token) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", url, false); // false for synchronous request

    if (token) {
        xmlHttp.setRequestHeader("token", token);
    }

    xmlHttp.send(null);
    return xmlHttp.responseText;
}
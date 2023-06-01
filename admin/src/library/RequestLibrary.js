export function httpGet(url) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", url, false); // false for synchronous request

    xmlHttp.send();
    return xmlHttp.responseText;
}

export function httpPost(url, body) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", url, false); // false for synchronous request
    console.log(body)
    xmlHttp.send(JSON.stringify(body));
    return xmlHttp.responseText;
}

export function httpPut(url, body) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("PUT", url, false); // false for synchronous request
    xmlHttp.send(JSON.stringify(body));
    return xmlHttp.responseText;
}

export function httpDelete(url) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("DELETE", url, false); // false for synchronous request

    xmlHttp.send();
    return xmlHttp.responseText;
}
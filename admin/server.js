'use strict';

const express = require('express');

const port = 9093;
const host = '0.0.0.0';

const app = express();

const serverHost = process.env.SERVER_HOST || '127.0.0.1';


app.use('/resources', express.static(__dirname + 'src/resources'));
app.use('/', express.static(__dirname + '/src'));

app.get('/', (req, res) => {
    res.sendFile('src/index', {serverHost: serverHost});
    //res.sendFile(__dirname + '/src/index.html');
});

app.listen(port, host, () => {
    console.log(`Running on http://${host}:${port}`);
});
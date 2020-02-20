const express = require('express');
const rp = require('request-promise');

const HOST = 3000

async function request () {
  console.log("let't go")
  await rp('http://localhost:9999/sleep')
  console.log('done')
}

const app = express();

app.get('/cpu', function (req, res) {
    let x = 0
    while (x <= 100000){
        console.log(x)
        x++
    }
    res.send('Hello World!');
});

app.get('/io', async function (req, res) {
    await request()
    res.send('Hello World!');
});

app.listen(HOST, function () {
  console.log(`Node.js sever listening on port ${HOST}!`);
});
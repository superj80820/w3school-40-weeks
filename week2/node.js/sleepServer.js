const express = require('express');
const app = express();

const HOST = 9999

app.get('/sleep', async function (req, res) {
    await new Promise(function (resolve, reject) {
      console.log("let's sleep")
      setTimeout(function () {
        console.log('sleep done')
        resolve()
    }, 5000)
    })
    res.sendStatus(200)
});

app.listen(HOST, function () {
  console.log(`SleepServer listening on port ${HOST}!`);
});
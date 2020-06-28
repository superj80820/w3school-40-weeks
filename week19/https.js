const fs = require('fs')
const https = require('https')
const express = require('express')
const app = express()
const privateKey  = fs.readFileSync('./privkey.pem')
const certificate = fs.readFileSync('./cert.pem')
const credentials = { key: privateKey, cert: certificate }

app.get('/', (res, req) => {
  req.send('Hello World!')
})

const httpsServer = https.createServer(credentials, app)
httpsServer.listen(8090)
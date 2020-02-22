# 測試方案

## 前置

1. `npm install`
2. Node server: `node node.js/server.js`
3. Sleep server: `node node.js/sleepServer.js`，這是模擬耗時IO的server

## IO bound測試

`node test/testIO.js`

## CPU bound測試

`node test/testCPU.js`
# 測試方案

## 前置

Terminal分別啟動兩個server，分別是

1. Node server: 執行`node node.js/server.js`來啟動
2. Sleep server: 執行`node node.js/sleepServer.js`來啟動，這是模擬耗時IO的server

## IO bound測試

執行`node test/testIO.js`

## CPU bound測試

執行`node test/testCPU.js`
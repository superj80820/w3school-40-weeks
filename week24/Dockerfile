# Server執行的環境，我們選用JDK 8
FROM openjdk:8

# Server所需的程式檔與配置，此指令是把week24資料夾所有的檔案都複製到Docker Image裡的/usr/src/myapp資料夾
COPY . /usr/src/myapp

# 說明Minecraft Server在Docker裡運行時的目錄
WORKDIR /usr/src/myapp

# Server執行的方式，我這邊就是把官網的執行指令複製下來而已
CMD ["java", "-Xmx1024M", "-Xms1024M", "-jar", "server.jar", "nogui"]
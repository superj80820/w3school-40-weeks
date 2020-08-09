FROM openjdk:8
COPY . /usr/src/myapp
WORKDIR /usr/src/myapp

CMD ["java", "-Xmx1024M", "-Xms1024M", "-jar", "server.jar", "nogui"]
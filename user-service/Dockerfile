FROM openjdk:17
VOLUME /tmp
ADD target/user-service.jar user-service.jar
EXPOSE 8133
ENTRYPOINT ["java","-jar","user-service.jar"]
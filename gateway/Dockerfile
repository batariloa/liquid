FROM openjdk:17
VOLUME /tmp
ADD target/apigateway.jar apigateway.jar
EXPOSE 9000
ENTRYPOINT ["java","-jar","apigateway.jar"]
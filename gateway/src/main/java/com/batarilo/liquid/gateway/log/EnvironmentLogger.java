package com.batarilo.liquid.gateway.log;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.stereotype.Component;

@Component
public class EnvironmentLogger implements ApplicationRunner {

    @Value("${STREAMING_SERVICE_URL:http://default-streaming-service-url}")
    private String streamingServiceUrl;

    @Value("${MEDIA_SERVICE_URL:http://default-media-service-url}")
    private String mediaServiceUrl;

    @Value("${SEARCH_SERVICE_URL:http://default-search-service-url}")
    private String searchServiceUrl;

    @Override
    public void run(ApplicationArguments args) {
        System.out.println("=== Environment Variables ===");
        System.out.println("STREAMING_SERVICE_URL: " + streamingServiceUrl);
        System.out.println("MEDIA_SERVICE_URL: " + mediaServiceUrl);
        System.out.println("SEARCH_SERVICE_URL: " + searchServiceUrl);
        System.out.println("=============================");
    }
}

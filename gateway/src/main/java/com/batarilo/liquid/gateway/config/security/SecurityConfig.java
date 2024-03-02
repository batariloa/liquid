package com.batarilo.liquid.gateway.config.security;


import lombok.RequiredArgsConstructor;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.web.server.SecurityWebFiltersOrder;
import org.springframework.security.config.web.server.ServerHttpSecurity;
import org.springframework.security.web.server.SecurityWebFilterChain;

@Configuration
@RequiredArgsConstructor
public class SecurityConfig
{

    private final JwtAuthenticationFilter authenticationFilter;

    @Bean
    SecurityWebFilterChain springWebFilterChain(ServerHttpSecurity http)
    {
        // Disable default security.
        http.httpBasic(ServerHttpSecurity.HttpBasicSpec::disable);
        http.formLogin(ServerHttpSecurity.FormLoginSpec::disable);
        http.csrf(ServerHttpSecurity.CsrfSpec::disable);
        http.logout(ServerHttpSecurity.LogoutSpec::disable);

        http.authorizeExchange(request ->
                request
                    .pathMatchers("/api/v1/auth/**").permitAll()
                    .anyExchange().authenticated())
            .addFilterAt(authenticationFilter,
                SecurityWebFiltersOrder.AUTHENTICATION);

        return http.build();
    }
}

package com.batarilo.liquid.gateway.service;

import io.jsonwebtoken.Jws;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.security.Key;
import java.util.Base64;

@Service
public class JwtService
{
    @Value("${security.jwt.public-key}")
    private String publicKeyString;

    public boolean validateToken(String token) {
        try {
            byte[] keyBytes = Base64.getDecoder().decode(publicKeyString);
            Key publicKey = Keys.hmacShaKeyFor(keyBytes);

            Jws<?> claims = Jwts.parserBuilder().setSigningKey(publicKey).build().parseClaimsJws(token);

            return true;
        } catch (Exception e) {
            return false;
        }
    }
}

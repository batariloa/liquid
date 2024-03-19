package com.batarilo.liquid.gateway.service;

import io.jsonwebtoken.Jws;
import io.jsonwebtoken.Jwts;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.security.KeyFactory;
import java.security.PublicKey;
import java.security.spec.X509EncodedKeySpec;
import java.util.Base64;

@Service
public class JwtService
{
    @Value("${security.jwt.public-key}")
    private String publicKeyString;

    public boolean validateToken(String token) {
        try {
            byte[] keyBytes = Base64.getDecoder().decode(publicKeyString);

            X509EncodedKeySpec keySpec = new X509EncodedKeySpec(keyBytes);
            KeyFactory keyFactory = KeyFactory.getInstance("RSA");
            PublicKey publicKey = keyFactory.generatePublic(keySpec);

            Jws<?> claims = Jwts.parser()
                .verifyWith(publicKey)
                .build().parseSignedClaims(token);

            return true;
        } catch (Exception e) {
            return false;
        }
    }
}

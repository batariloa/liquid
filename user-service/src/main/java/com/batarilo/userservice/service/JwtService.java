package com.batarilo.userservice.service;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import org.apache.tomcat.util.codec.binary.Base64;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Service;

import java.security.KeyFactory;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.spec.InvalidKeySpecException;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;
import java.time.Instant;
import java.time.temporal.ChronoUnit;
import java.util.Date;
import java.util.function.Function;

@Service
public class JwtService
{
    @Value("${security.jwt.public-key}")
    private String publicKeyString;

    @Value("${security.jwt.private-key}")
    private String privateKeyString;

    @Value("${security.jwt.expiration-time}")
    private long jwtExpiration;

    public String extractUsername(String token)
    {
        return extractClaim(token, Claims::getSubject);
    }

    public <T> T extractClaim(String token, Function<Claims, T> claimsResolver)
    {
        final Claims claims;
        try {
            claims = extractAllClaims(token);
        } catch (NoSuchAlgorithmException | InvalidKeySpecException e) {
            throw new RuntimeException(e);
        }
        return claimsResolver.apply(claims);
    }

    public String generateToken(UserDetails userDetails)
    {
        return buildToken(userDetails);
    }

    public long getExpirationTime()
    {
        return jwtExpiration;
    }

    private String buildToken(
        UserDetails userDetails
    )
    {
        PrivateKey privateKey = getPrivateKey();

        Instant now = Instant.now();

        return Jwts.builder()
            .setSubject(userDetails.getUsername())
            .setAudience("account-d.docusign.com")
            .claim("scope", "signature impersonation")
            .setIssuedAt(Date.from(now))
            .setExpiration(Date.from(now.plus(5l, ChronoUnit.MINUTES)))
            .signWith(privateKey)
            .compact();
    }

    public boolean isTokenValid(String token, UserDetails userDetails)
    {
        final String username = extractUsername(token);
        return (username.equals(userDetails.getUsername())) && !isTokenExpired(token);
    }

    private boolean isTokenExpired(String token)
    {
        return extractExpiration(token).before(new Date());
    }

    private Date extractExpiration(String token)
    {
        return extractClaim(token, Claims::getExpiration);
    }

    private Claims extractAllClaims(String token) throws NoSuchAlgorithmException, InvalidKeySpecException
    {
        return Jwts
            .parserBuilder()
            .setSigningKey(getPrivateKey())
            .build()
            .parseClaimsJws(token)
            .getBody();
    }

    private PrivateKey getPrivateKey()
    {

        String rsaPrivateKey = privateKeyString.replace("-----BEGIN PRIVATE KEY-----", "");
        rsaPrivateKey = rsaPrivateKey.replace("-----END PRIVATE KEY-----", "");

        PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(java.util.Base64.getDecoder().decode(rsaPrivateKey));
        KeyFactory kf = null;
        try {
            kf = KeyFactory.getInstance("RSA");
        } catch (NoSuchAlgorithmException e) {
            throw new RuntimeException(e);
        }

        try {
            return kf.generatePrivate(keySpec);
        } catch (InvalidKeySpecException e) {
            throw new RuntimeException(e);
        }
    }
}

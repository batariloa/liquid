package auth

import (
	"StorageService/internal/apierror"
	"crypto/rsa"
	"encoding/base64"
	"encoding/pem"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var publicKey *rsa.PublicKey

type CustomClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func init() {
	publicKeyBase64 := `MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEApWwDqgRTVbB9o4RiL4CTQnMszmhV1DXaKUFGPyxtgPqlsvRvpaaP4BJH6FTjQcMpo5rRLdSaF+9DDqMNfWzQvd7wcMtqIrJbep281nHXMCegdJ4dbRpF3+lHq4LRWyWl2A8F0OJZRFyCAvJrqhCSA45tLbIDiyT0Xb/UeEUbye+7pcGMnX/pIabQzLv4XNfI9tlCxPlb/xKIYVz+EAN6/d6SdmaqODB8ZdB1XObpnORy3KzY49wqMeeebykJpbXNEIrm6ROu8KSgIAvuf8mzQ80QjlcCOcri+K2gSl0I+5WBDRu59KnUZ1horyCpneXitoQG2ak8nK3ECYKGAcEgYjy+JIo6mOe6TloN9Ev872xk00APRW2UWV0pfMrVXVIqnKefiBrSvWC95yRQkYELjZ5ELLxQyyxe3fvNtqIcp+H94K7PGbKmK3rT5h436pOpQUPKfRmdBQqJX1wMZJ9q+5H6jx/VTulskBxZqZqOfED4jdWaXhKQE3ZuPuVHhnlZAgMBAAE=`

	// Decode the Base64-encoded RSA public key string
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		log.Println("Error decoding Base64 string:", err)
		return
	}

	// Create a PEM block containing the RSA public key
	pemBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	// Parse the PEM block to get the RSA public key
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(pem.EncodeToMemory(pemBlock))
	if err != nil {
		log.Println("Error parsing RSA public key:", err)
		return
	}
}

func ExtractUserIdFromToken(r *http.Request) (int, error) {
	tokenString := extractTokenFromRequest(r)
	if tokenString == "" {
		return 0, apierror.NewNotFoundError("No token found")
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil // Use the parsed RSA public key
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return 0, apierror.NewBadRequestError("Invalid claims")
	}

	return claims.UserId, nil
}

func extractTokenFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if token != "" && strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ")
	}

	return ""
}

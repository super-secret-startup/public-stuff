package main

import (
    "fmt"
    "time"
    "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userId string) string {
    claims := jwt.StandardClaims{
        Issuer: "server",
        Subject: userId,
        ExpiresAt: time.Now().Unix() + 3600,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
    tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

    return tokenString
}

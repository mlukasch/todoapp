package handlers

import "github.com/dgrijalva/jwt-go"
import "time"

type Claim struct {
	firstName string
	lastName  string
	email     string
	jwt.StandardClaims
}

func createClaims(firstName, lastName, email string) *Claim {
	claims := Claim{
		firstName,
		lastName,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).UTC().Unix(),
			Issuer:    "localhost:8080",
		},
	}
	return &claims
}

func createSignedToken(firstName, lastName, email string) string {
	claims := createClaims(firstName, lastName, email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("secret"))
	return signedToken
}

func LoginHandler(firstName, lastName, email string) {

}

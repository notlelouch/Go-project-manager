package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func WithJWTAuth(handlerFunc http.HandlerFunc, store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token from the Request
		tokenString := GetTokenFromRequest(r)
		// validate the token
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Println("Failed to authenticate token")
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("Failed to authenticate token")
			permissionDenied(w)
			return

		}

		// get the userID from the user
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(string)

		_, err = store.GetUserByID(userID)
		if err != nil {
			log.Println("Failed to get user")
			permissionDenied(w)
			return
		}

		// call the handlerFunc and continue to the endpoint if the token is Valid
		handlerFunc(w, r)
	}
}

func CreateJWT(secret []byte, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(time.Hour * 24 * 120).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusUnauthorized, ErrorResponse{
		Error: fmt.Errorf("permission denied").Error(),
	})
}

package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/NotFlexRestAPI/models"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(RandStringBytesRmndr(128))
var tokenName = "token"

type Claims struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"usertype"`
	jwt.StandardClaims
}

func RandStringBytesRmndr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	println("JWT Key :" + string(b))
	return string(b)
}

func generateToken(w http.ResponseWriter, id string, name string, userType int) {
	// Expire in 1 day
	tokenExpiryTime := time.Now().Add(60 * 24 * time.Minute)

	claims := &Claims{
		ID:       id,
		Name:     name,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}

func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken, id := validateUserToken(w, r, accessType)
		if id == "" {
			var response models.Response
			response.Status = 401
			response.Message = "Unauthorized Access"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else if !isValidToken {
			var response models.Response
			response.Status = 403
			response.Message = "Forbidden Access"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(w http.ResponseWriter, r *http.Request, accessType int) (bool, string) {
	isAccessTokenValid, id, _, userType := validateUserTokenFromCookies(r)
	if isAccessTokenValid {
		isUserValid := userType == accessType
		if accessType == 0 {
			isUserValid = true
		}
		if isUserValid {
			fmt.Print("ID User trigger : ", id)
			return true, id
		}
	}
	return false, id
}

func validateUserTokenFromCookies(r *http.Request) (bool, string, string, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.UserType
		}
	}
	return false, "", "", -1
}

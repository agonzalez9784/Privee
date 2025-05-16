package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

type Token struct{
	UserID string
	Exp time.Time
}
func ParseToken(tokenString string) Token{
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Error("Oops! Something went wrong!")
		}
	})

	if err != nil {
		
	}

 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID, ok := claims["userID"].(string)
        if !ok {
            return "", time.Time{}, fmt.Errorf("userID not found in token")
        }

        expFloat, ok := claims["exp"].(float64)
        if !ok {
            return "", time.Time{}, fmt.Errorf("exp not found or invalid in token")
        }

        expirationTime := time.Unix(int64(expFloat), 0)
	}

	parsedToken := Token{
		UserID: userID,
		Exp: expFloat
	}
	
	return parsedToken
}

func CreateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userID,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Println("Problem generating token...")
		return "", err
	}
	fmt.Println("The token was generating, but no output")
	fmt.Println(tokenString)
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

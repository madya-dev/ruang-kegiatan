package helper

import (
	"fmt"
	"madyasantosa/ruangkegiatan/config"
	"madyasantosa/ruangkegiatan/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(t *dto.Token) (string, error) {
	expireTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["username"] = t.Username
	claims["role"] = t.Role
	claims["exp"] = expireTime
	claims["iat"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(config.InitConfig().JWTSecret))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

func ExtractToken(tokenString string) (*dto.Token, error) {

	type CustomClaims struct {
		Username string `json:"username"`
		Role     string `json:"role"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.InitConfig().JWTSecret), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		extractedToken := &dto.Token{
			Username: claims.Username,
			Role:     claims.Role,
			Exp:      time.Unix(claims.ExpiresAt.Time.Unix(), 0),
			Iat:      time.Unix(claims.IssuedAt.Time.Unix(), 0),
		}
		return extractedToken, nil
	}

	return nil, fmt.Errorf("Invalid token")

}

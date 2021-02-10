package security

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey string
)

const (
	token_expire = (24 * time.Hour) * 5
	userIdClaims = "user_id"
	expireClaims = "expire"
)

func init() {
	secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		panic("secretKey is undefined")
	}
}

func CreateToken(id uint64) (string, error) {
	claims := jwt.MapClaims{}

	claims[userIdClaims] = id
	claims[expireClaims] = time.Now().UTC().Add(token_expire).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte(secretKey))

	return result, err

}

func isExpired(expireTime int64) bool {
	return expireTime < time.Now().UTC().Unix()
}

func ParseToken(rawToken string) (uint64, error) {
	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return 0, errors.New("{In ExtractTokenID} This should not happen")
	}
	userID, uintParseErr := strconv.ParseUint(fmt.Sprintf("%.0f", claims[userIdClaims]), 10, 64)
	expire, expireParseErr := strconv.ParseInt(fmt.Sprintf("%.0f", claims[expireClaims]), 10, 64)

	if isExpired(expire) {
		return 0, errors.New("Token is expired")
	}

	if uintParseErr != nil {
		return 0, uintParseErr
	}
	if expireParseErr != nil {
		return 0, expireParseErr
	}
	return uint64(userID), nil
}

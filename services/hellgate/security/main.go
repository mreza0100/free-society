package security

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

const (
	WriterKeyCtx  = "writter"
	RequestKeyCtx = "request"
	UserIdKeyCtx  = "user_id"
	CookieName    = "auth"
)

func SetToken(ctx context.Context, userId uint64) error {
	token, err := CreateToken(userId)
	if err != nil {
		return err
	}

	http.SetCookie(*GetWriter(ctx), &http.Cookie{
		Name:     CookieName,
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(token_expire),
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
	return nil
}

func DeleteToken(ctx context.Context) {
	http.SetCookie(*GetWriter(ctx), &http.Cookie{
		Name:     CookieName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
}

func extractUserId(req *http.Request) (uint64, error) {
	c, err := req.Cookie(CookieName)
	if err != nil {
		return 0, errors.New("There is no token in cookies")
	}

	userId, err := ParseToken(c.Value)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func GetWriter(ctx context.Context) *http.ResponseWriter {
	return ctx.Value(WriterKeyCtx).(*http.ResponseWriter)
}
func GetRequest(ctx context.Context) *http.Request {
	return ctx.Value(RequestKeyCtx).(*http.Request)
}
func GetUserId(ctx context.Context) uint64 {
	data, ok := ctx.Value(UserIdKeyCtx).(uint64)
	if !ok {
		log.Fatal("in GetUserId: ", "ok was false")
	}

	return data
}

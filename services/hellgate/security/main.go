package security

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	WRITE_KEY_CTX      = "writter"
	REQUEST_KEY_CTX    = "request"
	USER_ID_KEY_CTX    = "user_id"
	COOKIE_NAME        = "auth"
	COOKIE_EXPIRE_TIME = time.Hour * 24 * 7
)

func SetToken(ctx context.Context, token string) {
	http.SetCookie(*GetWriter(ctx), &http.Cookie{
		Name:     COOKIE_NAME,
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(COOKIE_EXPIRE_TIME),
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
}

func DeleteToken(ctx context.Context) {
	http.SetCookie(*GetWriter(ctx), &http.Cookie{
		Name:     COOKIE_NAME,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
}

func GetToken(ctx context.Context) string {
	req := GetRequest(ctx)

	c, _ := req.Cookie(COOKIE_NAME)
	return c.Value
}

func GetWriter(ctx context.Context) *http.ResponseWriter {
	return ctx.Value(WRITE_KEY_CTX).(*http.ResponseWriter)
}
func GetRequest(ctx context.Context) *http.Request {
	return ctx.Value(REQUEST_KEY_CTX).(*http.Request)
}
func GetUserId(ctx context.Context) uint64 {
	data, ok := ctx.Value(USER_ID_KEY_CTX).(uint64)
	if !ok {
		log.Fatal("in GetUserId: ", "ok was false")
	}

	return data
}
func GetOptionalId(ctx context.Context) (uint64, bool) {
	data, ok := ctx.Value(USER_ID_KEY_CTX).(uint64)
	return data, ok
}

func GetUserAgent(ctx context.Context) string {
	return GetRequest(ctx).Header.Get("user-agent")
}

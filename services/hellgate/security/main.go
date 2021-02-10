package security

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	cookieAccessKeyCtx = "user_id"
	cookieName         = "auth"
)

type CookieAccess struct {
	Writer      http.ResponseWriter
	NotLoginErr error
	UserId      uint64
	IsLoggedIn  bool
}

func (this *CookieAccess) SetToken(token string) {
	http.SetCookie(this.Writer, &http.Cookie{
		Name:     cookieName,
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(token_expire),
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
}

func (this *CookieAccess) DeleteToken() {
	http.SetCookie(this.Writer, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
}

func extractUserId(ctx *gin.Context) (uint64, error) {
	c, err := ctx.Request.Cookie(cookieName)
	if err != nil {
		return 0, errors.New("There is no token in cookies")
	}

	userId, err := ParseToken(c.Value)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func setValInCtx(ctx *gin.Context, val interface{}) {
	newCtx := context.WithValue(ctx.Request.Context(), cookieAccessKeyCtx, val)
	ctx.Request = ctx.Request.WithContext(newCtx)
}

var notLoginErr = errors.New("You are not logged in")

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookieA := CookieAccess{
			Writer:      ctx.Writer,
			NotLoginErr: notLoginErr,
			UserId:      0,
			IsLoggedIn:  false,
		}

		setValInCtx(ctx, &cookieA)

		userId, err := extractUserId(ctx)
		if err != nil {
			ctx.Next()
			return
		}

		cookieA.UserId = userId
		cookieA.IsLoggedIn = true

		ctx.Next()
	}
}

func GetCookieAccess(ctx context.Context) *CookieAccess {
	return ctx.Value(cookieAccessKeyCtx).(*CookieAccess)
}

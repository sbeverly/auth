package cookies

import (
	"net/http"
	"time"

	"github.com/sbeverly/auth/internal/config"
)

var cookieConf config.CookieConfig

func init() {
	cookieConf = config.GetConfig().Cookie
}

func GenerateLoginCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    token,
		Domain:   cookieConf.Domain,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour)}
}

func GenerateLogoutCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    "",
		Domain:   cookieConf.Domain,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(-1 * time.Hour)}
}

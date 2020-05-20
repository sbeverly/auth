package cookies

import (
	"net/http"

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
		MaxAge:   1 * 60 * 60}
}

func GenerateLogoutCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    "",
		Domain:   cookieConf.Domain,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   0}
}

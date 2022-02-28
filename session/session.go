package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifeTime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	minutes, err := strconv.Atoi(c.CookieLifeTime)
	if err != nil {
		minutes = 60
	}

	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	}

	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}

	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.SameSite = http.SameSiteLaxMode // ToRead more

	switch strings.ToLower(c.SessionType) {
	case "redis":
	case "mysql", "mariadb":
	case "postgres", "postgresql":
	default:
	}
	return session
}

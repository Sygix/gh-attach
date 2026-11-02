package browserprovider

import (
	"context"
	"net/http"

	"github.com/Sygix/gh-attach/internal/cookies"
)

type CookieSet struct {
	Profile   string
	Cookies   []*http.Cookie
	StorePath string
}

type Backend interface {
	Name() string
	Load(ctx context.Context, host string, source cookies.Source) ([]CookieSet, error)
}

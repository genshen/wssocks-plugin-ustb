package webview

import "net/http"

// WebviewAuth is an interface for passing data from UI
type WebviewAuth interface {
	// GetCookie returns the cookies after auth
	GetCookie(client *http.Client, loginUrl string) ([]*http.Cookie, error)
	// WaitAuthFinished wait auth finish. After finished, the websocket connection can be established.
	WaitAuthFinished() error
}

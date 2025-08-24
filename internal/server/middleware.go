package server

import "net/http"

func (app *AppConfig) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip       = r.RemoteAddr
			protocol = r.Proto
			url      = r.RequestURI
		)

		app.Logger.Info("received request", "ip", ip, "protocol", protocol, "url", url)

		next.ServeHTTP(w, r)
	})
}

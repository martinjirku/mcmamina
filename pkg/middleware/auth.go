package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"jirku.sk/mcmamina/pkg/models"
)

const SessionName = "session"

type UserCookie int

const UserCookieKey UserCookie = 0

func AuthMiddleware(store sessions.Store) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			session, err := store.Get(r, SessionName)
			if err != nil || session.Values["user"] != nil {
				if user, ok := session.Values["user"].(models.UserLogin); ok {
					ctx = context.WithValue(ctx, UserCookieKey, user)
				}
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetUser(r)
		if user == nil {
			// TODO: handle redirect back to original page
			http.Redirect(w, r, "/prihlasenie", http.StatusFound)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func StoreUser(w http.ResponseWriter, r *http.Request, user *models.UserLogin, store sessions.Store) error {
	session := sessions.NewSession(store, SessionName)
	if user == nil {
		session.Values["user"] = nil
	} else {
		session.Values["user"] = &user
	}

	session.Options.MaxAge = 60 * 60 * 24 * 10
	session.Options.HttpOnly = true
	session.Options.Path = "/"
	session.Save(r, w)
	return nil
}

func GetUser(r *http.Request) *models.UserLogin {
	result := r.Context().Value(UserCookieKey)
	if result == nil {
		return nil
	} else if result, ok := result.(models.UserLogin); ok {
		return &result
	}
	return nil
}

func BasicAuthMiddleware(username, password string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != username || pass != password {
			// If the authentication fails, return an unauthorized status and a realm
			w.Header().Set("WWW-Authenticate", `Basic realm="metrics"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		// If authentication succeeds, pass the request to the original handler
		next.ServeHTTP(w, r)
	})
}

package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type JSONRequestURL struct {
	LongURL string `json:"longurl"`
}

func isValidURL(token string) bool {
	_, err := url.ParseRequestURI(token)
	if err != nil {
		return false
	}

	u, err := url.Parse(token)
	if err != nil || u.Host == "" {
		return false
	}

	return true
}

func (s *HTTPService) dataValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			next.ServeHTTP(w, r)
		}

		bodyData := JSONRequestURL{}

		if err = json.Unmarshal(body, &bodyData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			next.ServeHTTP(w, r)
		}

		if ok := isValidURL(bodyData.LongURL); !ok {
			w.WriteHeader(http.StatusBadRequest)
			next.ServeHTTP(w, r)
		}

		ctx := context.WithValue(r.Context(), "LongURL", bodyData.LongURL)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

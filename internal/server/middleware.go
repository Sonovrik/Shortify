package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"

type JsonRequestUrl struct {
	LongUrl string `json:"longurl"`
}

func isValidUrl(token string) bool {
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

func (s *HttpService) dataValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			next.ServeHTTP(w, r)
		}

		bodyData := JsonRequestUrl{}

		if err = json.Unmarshal(body, &bodyData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			next.ServeHTTP(w, r)
		}

		if ok := isValidUrl(bodyData.LongUrl); !ok {
			w.WriteHeader(http.StatusBadRequest)
			next.ServeHTTP(w, r)
		}

		ctx := context.WithValue(r.Context(), "Long url", bodyData.LongUrl)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

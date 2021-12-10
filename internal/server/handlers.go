package server

import "net/http"

func (s *HTTPService) HandleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		longURL := ctx.Value("Long url")
		if longURL == nil {
			// bad request or internal request
			return
		}
	}
}

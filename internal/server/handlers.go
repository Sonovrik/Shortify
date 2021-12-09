package server

import "net/http"

func (s *HttpService) HandleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		longUrl := ctx.Value("Long url")
		if longUrl == nil {
			// bad request or internal request
			return
		}

		//... next steps

	}
}

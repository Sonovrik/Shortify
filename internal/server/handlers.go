package server

import "net/http"

func (s *HttpService) HandleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

package server

import (
	"net/http"
	"net/url"
	"time"

	"rubrumcreation.com/go-paper-scissor/services"
	"rubrumcreation.com/go-paper-scissor/util"
)


func makeHandler(fn func(http.ResponseWriter, *http.Request, url.Values)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		client := util.GetIpFromRequest(r)

		query := r.URL.Query()

		fn(w, r, query)
		services.LogApiInfo(r.URL.Path, client, "API request finished", time.Since(start).Abs().Microseconds())
	}
}

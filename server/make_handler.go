package server

import (
	"net/http"
	"net/url"
	"regexp"
	"time"

	"rubrumcreation.com/go-paper-scissor/services"
	"rubrumcreation.com/go-paper-scissor/util"
)

var validPath = regexp.MustCompile("^/(simulation|play)/(random+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, url.Values)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		client := util.GetIpFromRequest(r)
		m := validPath.FindStringSubmatch(r.URL.Path)
		query := r.URL.Query()
		if m == nil {
			services.LogApiInfo(r.URL.Path, client, "Unkown route reqested by client", time.Since(start).Abs().Microseconds())
			http.NotFound(w, r)
			return
		}

		fn(w, r, query)
		services.LogApiInfo(r.URL.Path, client, "API request finished", time.Since(start).Abs().Microseconds())
	}
}

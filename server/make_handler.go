package server

import (
	"net/http"
	"regexp"
	"time"

	"rubrumcreation.com/go-paper-scissor/services"
	"rubrumcreation.com/go-paper-scissor/util"
)

var validPath = regexp.MustCompile("^/(random|simulate)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		start := time.Now()
		client := util.GetIpFromRequest(r)
		fn(w, r, m[2])
		services.LogApiInfo(r.URL.Path, client, "API request finished", time.Since(start).Abs().Microseconds())
	}
}

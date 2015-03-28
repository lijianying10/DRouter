package Router

import (
	"DRouter/conf"
	"DRouter/logD"
	"net/http"
	"net/http/httputil"
)

// Bindings binding service router
func Bindings() {
	for rout, dist := range conf.Static {
		http.Handle(rout, http.StripPrefix(rout, http.FileServer(http.Dir(dist.(string)))))
	}

	for rout, dist := range conf.Router {
		http.HandleFunc(rout, func(w http.ResponseWriter, r *http.Request) {
			logD.Logger.Print(r.URL, " -> ")
			director := func(req *http.Request) {
				req = r
				req.URL.Scheme = "http"
				req.URL.Host = dist.(string)
				logD.Logger.Print(req.URL, "\n")
			}
			proxy := &httputil.ReverseProxy{Director: director}
			proxy.ServeHTTP(w, r)
		})
	}
}

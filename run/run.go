package command

import "net/http"

var Run = func() {
	jakgo := NewJakGoController()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1", jakgo.JakGo)
	http.ListenAndServe(":5000", mux)
}

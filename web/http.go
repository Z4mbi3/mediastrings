package web

import (
	"net/http"
)

func RunWeb() {
	fs := http.FileServer(http.Dir("./web/public"))
	http.Handle("/", fs)
	http.ListenAndServe(":9999", nil)
}

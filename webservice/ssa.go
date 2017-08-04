package webservice

import (
	"net/http"

	"github.com/nickng/gospal/ssa/build"
)

func initSSA() {
	http.HandleFunc("/ssa", ssaHandler)
}

func ssaHandler(w http.ResponseWriter, req *http.Request) {
	info, err := build.FromReader(req.Body).Default().Build()
	req.Body.Close()
	if err != nil {
		NewErrInternal(err, "Cannot build SSA").Report(w)
		return
	}
	info.WriteTo(w)
}

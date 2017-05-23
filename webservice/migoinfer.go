package webservice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/nickng/gospal/lib/migoinfer"
	"github.com/nickng/gospal/ssa/build"
)

func migoinferHandler(w http.ResponseWriter, req *http.Request) {
	conf := build.FromReader(req.Body).Default()
	err := req.Body.Close()
	if err != nil {
		NewErrInternal(err, "Cannot initialise SSA").Report(w)
	}
	info, err := conf.Build()
	if err != nil {
		NewErrInternal(err, "Cannot build SSA").Report(w)
	}
	inferer := migoinfer.New(info, os.Stderr)
	var out bytes.Buffer
	inferer.SetOutput(&out)
	inferer.Analyse()
	reply := struct {
		MiGo string `json:"MiGo"`
		Time string `json:"time"`
	}{
		MiGo: out.String(),
		Time: "0",
	}
	json.NewEncoder(w).Encode(&reply)
}

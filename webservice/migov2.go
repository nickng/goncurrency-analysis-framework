package webservice

// MiGo inference module (v2).
// This uses the new gospal MiGo inference, with for-loop support.

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/nickng/gospal/migoinfer"
	"github.com/nickng/gospal/ssa/build"
)

func initMigoV2() {
	http.HandleFunc("/migo.v2", migoV2Handler)
}

func migoV2Handler(w http.ResponseWriter, req *http.Request) {
	conf := build.FromReader(req.Body).Default()
	err := req.Body.Close()
	if err != nil {
		NewErrInternal(err, "Cannot initialise SSA").Report(w)
		return
	}
	info, err := conf.Build()
	if err != nil {
		NewErrInternal(err, "Cannot build SSA").Report(w)
		return
	}
	inferer := migoinfer.New(info, os.Stderr)
	var out bytes.Buffer
	inferer.SetOutput(&out)
	inferer.Raw = false
	startTime := time.Now()
	inferer.Analyse()
	execTime := time.Now().Sub(startTime)
	reply := struct {
		MiGo string `json:"MiGo"`
		Time string `json:"time"`
	}{
		MiGo: out.String(),
		Time: execTime.String(),
	}
	json.NewEncoder(w).Encode(&reply)
}

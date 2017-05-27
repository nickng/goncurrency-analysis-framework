package webservice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/nickng/gospal/lib/migoinfer"
	"github.com/nickng/gospal/ssa/build"
)

func initMigoinfer() {
	http.HandleFunc("/infer", migoinferHandler)
}

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

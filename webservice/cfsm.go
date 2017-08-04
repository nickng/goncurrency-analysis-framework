package webservice

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nickng/dingo-hunter/cfsmextract"
	"github.com/nickng/dingo-hunter/cfsmextract/sesstype"
	"github.com/nickng/dingo-hunter/ssabuilder"
)

func initCFSMExtract() {
	http.HandleFunc("/cfsm", cfsmHandler)
}

func cfsmHandler(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		NewErrInternal(err, "Cannot read input Go source code").Report(w)
		return
	}
	req.Body.Close()
	conf, err := ssabuilder.NewConfigFromString(string(b))
	if err != nil {
		NewErrInternal(err, "Cannot initialise SSA").Report(w)
		return
	}
	ssainfo, err := conf.Build()
	if err != nil {
		NewErrInternal(err, "Cannot build SSA").Report(w)
		return
	}
	extract := cfsmextract.New(ssainfo, "extract", "/tmp")
	go extract.Run()

	select {
	case <-extract.Error:
		NewErrInternal(err, "CFSM extraction failed").Report(w)
		return
	case <-extract.Done:
		log.Println("CFSMs: analysis completed in", extract.Time)
	}
	cfsms := sesstype.NewCFSMs(extract.Session())
	bufCfsm := new(bytes.Buffer)
	cfsms.WriteTo(bufCfsm)
	dot := sesstype.NewGraphvizDot(extract.Session())
	bufDot := new(bytes.Buffer)
	dot.WriteTo(bufDot)
	reply := struct {
		CFSM string `json:"CFSM"`
		Dot  string `json:"dot"`
		Time string `json:"time"`
	}{
		CFSM: bufCfsm.String(),
		Dot:  bufDot.String(),
		Time: extract.Time.String(),
	}
	json.NewEncoder(w).Encode(&reply)
}

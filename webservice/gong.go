package webservice

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func initGong() {
	http.HandleFunc("/gong", gongHandler)
}

func gongHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Running Gong on snippet")
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		NewErrInternal(err, "Cannot read input MiGo types").Report(w)
		return
	}
	req.Body.Close()
	file, err := ioutil.TempFile(os.TempDir(), "gong")
	if err != nil {
		NewErrInternal(err, "Cannot create temp file for MiGo input").Report(w)
		return
	}
	defer os.Remove(file.Name())

	if _, err := file.Write(b); err != nil {
		NewErrInternal(err, "Cannot write to temp file for MiGo input").Report(w)
		return
	}
	if err := file.Close(); err != nil {
		NewErrInternal(err, "Cannot close temp file for MiGo input").Report(w)
		return
	}
	Gong, err := exec.LookPath("Gong")
	if err != nil {
		NewErrInternal(err, "Cannot find Gong executable (Check $PATH?)").Report(w)
		return
	}
	startTime := time.Now()
	out, err := exec.Command(Gong, file.Name()).CombinedOutput()
	if err != nil {
		log.Printf("Gong execution failed: %v\n", err)
	}
	execTime := time.Now().Sub(startTime)
	replacer := strings.NewReplacer("[92m", "<span style='color: #87ff87; font-weight: bold'>", "[91m", "<span style='color: #ff005f; font-weight: bold'>", "[0m", "</span>")
	reply := struct {
		Gong string `json:"Gong"`
		Time string `json:"time"`
	}{
		Gong: replacer.Replace(string(out)),
		Time: execTime.String(),
	}
	log.Println("Gong completed in", execTime.String())
	json.NewEncoder(w).Encode(&reply)
}

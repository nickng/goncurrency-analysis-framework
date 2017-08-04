package webservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

func initGodel() {
	http.HandleFunc("/godel", godelHandler)
	http.HandleFunc("/godelterm", godelTermHandler)
}

func godelHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Running Godel on snippet")
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		NewErrInternal(err, "Cannot read input MiGo types")
		return
	}
	req.Body.Close()
	file, err := ioutil.TempFile(os.TempDir(), "godel")
	os.Chdir(os.TempDir())
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
	Godel, err := exec.LookPath("docker")
	if err != nil {
		NewErrInternal(err, "Cannot find Godel executable (Check $PATH?)").Report(w)
		return
	}
	startTime := time.Now()

	file.Chdir()
	out, err := exec.Command(Godel, "run", "--rm", "-v", fmt.Sprintf("%s:/root", path.Dir(file.Name())), "nickng/godel:latest", "Godel", path.Base(file.Name())).CombinedOutput()
	if err != nil {
		log.Printf("Godel execution failed: %v\n", err)
	}
	execTime := time.Now().Sub(startTime)
	replacer := strings.NewReplacer("[92m", "<span style='color: #87ff87; font-weight: bold'>", "[91m", "<span style='color: #ff005f; font-weight: bold'>", "[0m", "</span>")
	reply := struct {
		Godel string `json:"Godel"`
		Time  string `json:"time"`
	}{
		Godel: replacer.Replace(string(out)),
		Time:  execTime.String(),
	}
	log.Println("Godel completed in", execTime.String())
	json.NewEncoder(w).Encode(&reply)
}

func godelTermHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Running Godel/Term on snippet")
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		NewErrInternal(err, "Cannot read input MiGo types")
		return
	}
	req.Body.Close()
	file, err := ioutil.TempFile(os.TempDir(), "godel")
	os.Chdir(os.TempDir())
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
	Godel, err := exec.LookPath("docker")
	if err != nil {
		NewErrInternal(err, "Cannot find Godel executable (Check $PATH?)").Report(w)
		return
	}
	startTime := time.Now()
	file.Chdir()
	out, err := exec.Command(Godel, "run", "--rm", "-v", fmt.Sprintf("%s:/root", path.Dir(file.Name())), "nickng/godel:latest", "Godel", "-T", path.Base(file.Name())).CombinedOutput()
	if err != nil {
		log.Printf("Godel execution failed: %v\n", err)
	}
	execTime := time.Now().Sub(startTime)
	replacer := strings.NewReplacer("[92m", "<span style='color: #87ff87; font-weight: bold'>", "[91m", "<span style='color: #ff005f; font-weight: bold'>", "[0m", "</span>")
	reply := struct {
		Godel string `json:"Godel"`
		Time  string `json:"time"`
	}{
		Godel: replacer.Replace(string(out)),
		Time:  execTime.String(),
	}
	log.Println("Godel/Term completed in", execTime.String())
	json.NewEncoder(w).Encode(&reply)
}

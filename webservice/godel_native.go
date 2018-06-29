// +build native

package webservice

import (
	"os"
	"os/exec"
)

func getGodelRunParams(f *os.File) ([]string, error) {
	Godel, err := exec.LookPath("Godel")
	if err != nil {
		return nil, err
	}
	return []string{Godel, f.Name()}, nil
}

func getGodelTermRunParams(f *os.File) ([]string, error) {
	Godel, err := exec.LookPath("Godel")
	if err != nil {
		return nil, err
	}
	return []string{Godel, "-T", f.Name()}, nil
}

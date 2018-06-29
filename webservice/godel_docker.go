// +build !native

package webservice

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func getGodelRunParams(f *os.File) ([]string, error) {
	docker, err := exec.LookPath("docker")
	if err != nil {
		return nil, err
	}
	// working directory of container is /root
	// mounts dirname of f as /root then run Godel with basename of f (i.e. /root/f)
	return []string{docker, "run", "--rm", "-v", fmt.Sprintf("%s:/root", path.Dir(f.Name())), "nickng/godel:latest", "Godel", path.Base(f.Name())}, nil
}

func getGodelTermRunParams(f *os.File) ([]string, error) {
	docker, err := exec.LookPath("docker")
	if err != nil {
		return nil, err
	}
	// working directory of container is /root
	// mounts dirname of f as /root then run Godel with basename of f (i.e. /root/f)
	return []string{docker, "run", "--rm", "-v", fmt.Sprintf("%s:/root", path.Dir(f.Name())), "nickng/godel:latest", "Godel", "-T", path.Base(f.Name())}, nil
}

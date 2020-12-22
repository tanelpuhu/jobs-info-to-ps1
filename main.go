package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

func getStdin() []byte {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error getting stdin: %v", err)
	}
	return data
}

func main() {
	input := strings.TrimSpace(string(getStdin()))
	var res []string
	for _, pidtx := range strings.Split(input, "\n") {
		if pidtx == "" {
			continue
		}
		pid, err := strconv.Atoi(pidtx)
		if err != nil {
			fmt.Printf("this pid is not numeric: %s (%v)\n", pidtx, err)
			continue
		} else if pid > 0 {
			p, err := ps.FindProcess(pid)
			if err != nil {
				fmt.Printf("error getting process %d info: %v\n", pid, err)
			}
			res = append(res, p.Executable())
		}
	}
	if len(res) > 0 {
		fmt.Printf(":%s", strings.Join(res, ","))
	}
}

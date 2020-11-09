package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("/bin/sh", "./start.sh")
    stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr) 
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

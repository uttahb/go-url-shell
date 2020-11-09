package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("/bin/sh", "./start.sh")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout
    if err := cmd.Run(); err != nil {
        fmt.Println( "Error:", err );
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

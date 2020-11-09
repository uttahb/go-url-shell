package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("in handler.....")
    cmd := exec.Command("/bin/sh", "./start.sh")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout
    // run command
    if err := cmd.Run(); err != nil {
        fmt.Println( "Error:", err )
        fmt.Fprintf(w, "err")
    }

    fmt.Fprintf(w, "Done\n")
}

func main() {
    http.HandleFunc("/build", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

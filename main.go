package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Ref string `json:"ref"`
	}
	fmt.Println("in handler.....")
	repo := r.URL.Query().Get("repo")
	r.ParseForm()

	var newr = Request{}
	payload := json.Unmarshal([]byte(r.Form["payload"][0]), &newr)
	fmt.Println(payload)

	if repo != "" {
		// ... process it, will be the first (only) if multiple were given
		// note: if they pass in like ?param1=&param2= param1 will also be "" :|
		file := "./start.sh"
		cmd := exec.Command("/bin/sh", file, repo)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
			fmt.Fprintf(w, "err")
		}

		fmt.Fprintf(w, "Done\n")
	}
	// run command

}

func main() {
	http.HandleFunc("/build", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

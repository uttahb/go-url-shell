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
	type Req struct {
		Ref string `json:"ref"`
	}
	fmt.Println("in handler.....")
	repo := r.URL.Query().Get("repo")
	r.ParseForm()
	var newr = Req{}
	fmt.Println(r.Form["payload"][0])
	payload := json.Unmarshal([]byte(r.Form["payload"][0]), &newr)
	fmt.Println(payload)
	fmt.Println(newr)
	if newr.Ref == "refs/heads/development" {
		fmt.Println("push to dev branch came, starting deploy")
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
	} else {
		fmt.Println("This is not the dev branch ", newr.Ref)
	}
	// run command

}

func main() {
	http.HandleFunc("/build", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

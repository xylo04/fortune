package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func execFortune() string {
	cmd := exec.Command("fortune")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

// fortune, the web server
func FortuneServer(w http.ResponseWriter, req *http.Request) {
	fortuneTxt := execFortune()
	body := fmt.Sprintf("<html><body><pre>\n%v\n</pre></body></html>", fortuneTxt)
	io.WriteString(w, body)
}

func main() {
	http.HandleFunc("/fortune", FortuneServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

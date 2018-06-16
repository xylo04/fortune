package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func execFortune() string {
	cmd := exec.Command("/usr/games/fortune")
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
	io.WriteString(w, `<html><head><title>Fortune</title></head>
    <body><pre>`)
	io.WriteString(w, fortuneTxt)
	io.WriteString(w, "</pre></body></html>")
}

func main() {
	port := flag.Int("port", 8080, "http port to listen on")
	flag.Parse()

	http.HandleFunc("/fortune", FortuneServer)
	addr := fmt.Sprintf(":%v", *port)
	log.Printf("Listening on %v, CTRL-C to quit...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

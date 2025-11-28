package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const INDEX = "/home/michele/Codice/daily-peanuts/index.html"
const GETSTRIP = "/home/michele/Codice/daily-peanuts/getStrip.js"
const ADDRESS = "127.0.0.1:8080"

func main() {
	var imgsrc string
	var out []byte
	out, _ = exec.Command("node", GETSTRIP).Output()
	imgsrc = strings.TrimSpace(string(out))
	fmt.Println("Retrieved image url:", imgsrc)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.Open(INDEX)
		defer f.Close()
		content, _ := io.ReadAll(f)
		content = []byte(strings.Replace(string(content), "{IMGSRC}", imgsrc, -1))
		w.Write(content)
	})
	fmt.Println("Serving at", ADDRESS)
	http.ListenAndServe(ADDRESS, nil)
}

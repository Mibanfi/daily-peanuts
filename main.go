package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const INDEX = "index.html"
const GETSTRIP = "getStrip.js"
const ADDRESS = "127.0.0.1:8080"

type saveData struct {
	page  []byte
	day   int
	month time.Month
	year  int
}

func (cache *saveData) update() {
	var page []byte
	var out []byte
	var dir, getstrip, index string

	day, month, year := time.Now().Date()

	if !(day == cache.day && month == cache.month && year == cache.year) {
		e, _ := os.Executable()
		dir = path.Dir(e)
		getstrip = filepath.Join(dir, GETSTRIP)
		index = filepath.Join(dir, INDEX)

		out, _ = exec.Command("node", getstrip).Output()
		imgsrc := strings.TrimSpace(string(out))
		fmt.Println("Scraped new image url from webpage:", imgsrc)

		f, _ := os.Open(index)
		page, _ = io.ReadAll(f)
		page = []byte(strings.Replace(string(page), "{IMGSRC}", imgsrc, -1))
		f.Close()
		fmt.Println("Built today's page")

		cache.page = page
		cache.day = day
		cache.month = month
		cache.year = year
	}

}

func main() {

	var cache saveData
	cache.update()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cache.update()
		fmt.Println("Serving page at", time.Now())
		w.Write(cache.page)
	})

	fmt.Println("Serving at", ADDRESS)
	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Closed")
	}
}

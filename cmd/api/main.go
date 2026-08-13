package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", r.URL, r.URL.Path)

	//fmt.Fprintf(w, "Hello, World! hey %s", r.URL.Path[1:])
	fmt.Println(r.URL.Path[1:])

	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}

	page, err := ioutil.ReadAll(res.Body)
	fmt.Fprint(w, string(page))
	fmt.Println(res.StatusCode)

	client := &http.Client{Timeout: time.Duration(10 * time.Second)}
	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	req.Header.Add("Accept-Encoding", "gzip")
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	var body []byte
	res.Body.Read(body)
	fmt.Println(body)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

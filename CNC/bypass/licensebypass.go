package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
  "log"
	"time"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	unescape, err := url.PathUnescape(string(all))
	if err != nil {
		return
	}
  w.Write([]byte(fmt.Sprintf("Valid|" + fmt.Sprintf("%d", time.Unix(0, time.Now().Unix()).Add(time.Duration(3000)*24)) + "")))
}

func main() {
	http.HandleFunc("/check", ApiHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 9999), nil); err != nil {
		log.Fatal(err)
	}
}

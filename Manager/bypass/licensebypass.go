package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Licensing request received, bypassed.")
	w.Write([]byte(fmt.Sprintf("Valid|" + fmt.Sprintf("%d", time.Unix(0, time.Now().Unix()).Add(time.Duration(3000)*24).Unix()))))
}

func main() {
	http.HandleFunc("/check", ApiHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 9999), nil); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/fendijatmiko/Context/log"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("localhost:8181", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printt(ctx, "server started")
	defer log.Printt(ctx, "server done")

	select {
	case <-time.After(time.Second * 3):
		fmt.Fprintf(w, "hayy from server")
	case <-ctx.Done():
		err := ctx.Err()
		log.Printt(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

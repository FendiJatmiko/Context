package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fendijatmiko/Context/log"
)

func main() {
	ctx := context.Background()
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8181", nil)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}
	req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}

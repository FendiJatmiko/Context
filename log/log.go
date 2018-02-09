package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

const IDkey = 56

func Printt(ctx context.Context, msg string) {
	id, ok := ctx.Value(IDkey).(int64)
	if !ok {
		log.Println("could not request ID in context")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, IDkey, id)
		f(w, r.WithContext(ctx))
	}
}

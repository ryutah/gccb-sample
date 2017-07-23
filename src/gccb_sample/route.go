package gccb_sample

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func InitRoute() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		log.Debugf(ctx, "This is sample")
		fmt.Fprintln(w, "Hello World!!!")
	})
}

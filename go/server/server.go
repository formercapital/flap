package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cockroachdb/pebble"
)

func Serve(port, libraryPath string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db, err := pebble.Open(libraryPath+"db", &pebble.Options{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		key := []byte("now")
		if err := db.Set(key, []byte(time.Now().String()), pebble.Sync); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		value, closer, err := db.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dbres := fmt.Sprintf("%s is %s", key, value)
		if err := closer.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := db.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Hello from Go\n\nport: %s\nfrom local DB:\n%s\n", port, dbres)
	})
	http.ListenAndServe(":"+port, mux)
}

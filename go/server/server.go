package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/cockroachdb/pebble"
)

func Serve(port, libraryPath string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dt := time.Now().String()
		f, err := os.Create(libraryPath + "test")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		f.WriteString("from local disk: " + time.Now().String())
		f.Close()
		f2, err := os.Open(libraryPath + "test")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		b := &bytes.Buffer{}
		_, err = io.Copy(b, f2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		f2.Close()
		db, err := pebble.Open(libraryPath+"demo", &pebble.Options{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		key := []byte("hello")
		if err := db.Set(key, []byte("world"), pebble.Sync); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		value, closer, err := db.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		dbres := fmt.Sprintf("%s %s\n", key, value)
		if err := closer.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := db.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "port: %s time: %s\n%s\n\n%s", port, dt, b.String(), dbres)
	})
	http.ListenAndServe(":"+port, mux)
}

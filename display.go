package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func file_response(fileName string, w http.ResponseWriter, r *http.Request) {
	fmt.Println(fileName)
	txt, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", txt)
}

func error_handler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		file_response("2048/404.html", w, r)
	}
}

func display(plan string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover_string := recover(); recover_string != nil {
				error_handler(w, r, http.StatusNotFound)
			}
		}()
		var fileName string
		if r.URL.Path == "/" {
			//file_response("2048/index.html", w, r)
			fileName = "2048/index.html"
		} else {
			fileName = "2048/" + r.URL.Path[1:]
		}
		if _, err := os.Stat(fileName); err == nil {
			// File exists
			http.ServeFile(w, r, fileName)
		} else {
			panic(fmt.Sprintf("File not found %s", fileName))
		}

	}
	http.HandleFunc("/", handler)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go http.ListenAndServe("localhost:8080", nil)
	open.Run("http://localhost:8080")
	wg.Wait()
}

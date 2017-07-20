package main

import (
	"flag"
	"log"
	"net/http"
)

var p = flag.String("p", "8080", "TCP port")
var d = flag.String("d", "/", "target directory")

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*d)))
	//http.HandleFunc("/view", makeHandler(getFileHandler))
	log.Fatal(http.ListenAndServe("localhost:"+*p, nil))
}

/*
func fetchFile(title string) ([]byte, error) {
	filename := title + ".txt"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func getFileHandler(w http.ResponseWriter, req *http.Request, title string) {
	title = *d
	substr := title[1:]
	b, err := fetchFile(substr)
	if err != nil {
		fmt.Fprint(w, "could not fetch file. File not found ", http.StatusNotFound)
		return
	}
	s := string(b)
	fmt.Fprint(w, s)
	return
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		title := req.URL.Path
		fn(w, req, title)
	}
}

//func hello(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "%v", "hello, Gopher! and welcome")
//}
*/

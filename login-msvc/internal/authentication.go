package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
func Connection(w http.ResponseWriter, req *http.Request) {

// Functions serving as handlers take a
// `http.ResponseWriter` and a `http.Request` as
// arguments. The response writer is used to fill in the
// HTTP response. Here our simple response is just
// "hello\n".
	fmt.Fprintf(w, "Connection: OK \n")
}

func Headers(w http.ResponseWriter, req *http.Request) {

// This handler does something a little more
// sophisticated by reading all the HTTP request
// headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
		fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func GET(w http.ResponseWriter, req *http.Request) {

	response, err := http.Get("http://golang.org/")
	
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}
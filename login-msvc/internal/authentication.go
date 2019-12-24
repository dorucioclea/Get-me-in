package internal

import (
	"encoding/json"
	"fmt"
	"github/Get-me-in/login-msvc/pkg/models"
	"net/http"
)

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
func APIInfo(w http.ResponseWriter, req *http.Request) {
	
	m := models.Message{"1.0", "GO", "1.13.5"}
	b, err := json.Marshal(m)

	if err != nil {
		fmt.Sprintf(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(b))
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

	/* EXAMPLE GET Request */
	/*
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
	*/
}
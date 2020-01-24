package internal

import "net/http"

// New returns an error that formats as the given text.
func CustomerError(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func HandleError(err error, w http.ResponseWriter){
	if err != nil {
		http.Error(w, err.Error(), http.StatusFailedDependency)
	}
}
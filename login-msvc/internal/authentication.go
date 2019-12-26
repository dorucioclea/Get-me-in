package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/Get-me-in/login-msvc/internal/security"
	"github/Get-me-in/login-msvc/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

func VerifyCredentials(w http.ResponseWriter, req *http.Request) {

	body, err2 := ioutil.ReadAll(req.Body)

	if err2 != nil{
		log.Fatal(err2)
	}

	resp, err := http.Post("http://localhost:8080/mock", "application/json" , bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {

		m := models.Message{"1.0",
			"GO",
			runtime.Version(),
			security.GenerateToken()}
		b, err := json.Marshal(m)

		if err != nil {
			fmt.Sprintf(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(b))	}

	w.WriteHeader(http.StatusUnauthorized)
}

func MockResponse(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)
}

package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/login-msvc/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

func VerifyCredentials(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil{
		log.Fatal(err)
	}

	resp, respErr := http.Post(configs.VERIFY_ACCOUNT, "application/json" , bytes.NewBuffer(body))

	if respErr != nil {
		log.Fatal(respErr)
	}

	if resp.StatusCode == 200 {

		m := Message{configs.API_VERSION,
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
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}
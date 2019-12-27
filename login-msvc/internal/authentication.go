package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/Get-me-in/login-msvc/configs"
	"github/Get-me-in/login-msvc/internal/security"
	"github/Get-me-in/login-msvc/pkg/models"
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

		m := models.Message{configs.API_VERSION,
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
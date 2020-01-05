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
	"time"
)

func VerifyCredentials(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil{
		log.Fatal(err)
	}

	resp, respErr := http.Post(configs.VERIFY_ACCOUNT ,"application/json" , bytes.NewBuffer(body))

	if respErr != nil {
		log.Fatal(respErr)
	}

	t := time.Now()
	e := t.Add(configs.EXPIRY * time.Minute)
	if resp.StatusCode == 200 {

		t := &security.TokenClaims{
			Issuer: configs.SERVICE_ID,
			Subject:    configs.SUBJECT,
			Audience:   req.Header.Get("Origin"),
			IssuedAt:   t.Unix(),
			Expiration: e.Unix(),
			NotBefore:  t.Unix(),
			Id:         "a",
		}
		fmt.Println(t)

		m := security.TokenResponse{
			AccessToken:  security.GenerateToken(t),
			TokenType:    configs.BEARER,
			ExpiresIn:    configs.EXPIRY,
			RefreshToken: "N/A",
		}

		b, err := json.Marshal(m)

		if err != nil {
			fmt.Sprintf(err.Error())
		}

		security.VerifyToken(m.AccessToken)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(b))
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func MockResponse(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}
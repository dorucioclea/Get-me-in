package internal

import (
	"../../pkg/security"
	"../configs"
	"bytes"
	"encoding/json"
	"fmt"
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

	resp, respErr := http.Post(configs.VERIFY_ACCOUNT, "application/json" , bytes.NewBuffer(body))

	if respErr != nil {
		log.Fatal(respErr)
	}

	t := time.Now()
	e := t.Add(360 * time.Minute).String()
	if resp.StatusCode == 200 {

		t := security.TokenClaims{
			Issuer: configs.SERVICE_ID,
			Subject:    configs.SUBJECT,
			Audience:   req.Header.Get("Origin"),
			IssuedAt:   t.String(),
			Expiration: e,
			NotBefore:  e,
			Id:         "a",}

		m := security.TokenResponse{
			security.GenerateToken(t),
			configs.BEARER,
			configs.EXPIRY,
			"N/A"}

		b, err := json.Marshal(m)

		if err != nil {
			fmt.Sprintf(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(b))
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func MockResponse(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}
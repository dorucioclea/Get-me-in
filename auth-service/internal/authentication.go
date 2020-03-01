package internal

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/auth-service/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/http_lib"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"io/ioutil"
	"net/http"
	"time"
)

func VerifyCredentials(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}


	//post, err := http.NewRequest("POST", configs.LOGIN_ENDPOINT, bytes.NewBuffer(body))
	//
	//post.Header.Set("Authorization", req.Header.Get("Authorization"))
	//
	//client := &http.Client{}
	//resp, err := client.Do(post)
	//
	//if err != nil {
	//	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	//}
	//
	//if resp.StatusCode != 200 {
	//	http.Error(w, http.StatusText(401), http.StatusUnauthorized)
	//}

	m := map[string]string{"Authorization": req.Header.Get("Authorization")}
	resp, err := http_lib.Post(configs.LOGIN_ENDPOINT, body, m)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	if resp.StatusCode != 200 {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
	}

	t := time.Now()
	e := t.Add(configs.EXPIRY * time.Minute)

	token := &security.TokenClaims{
		Issuer:     configs.SERVICE_ID,
		Subject:    configs.SUBJECT,
		Audience:   req.Header.Get("Origin"),
		IssuedAt:   t.Unix(),
		Expiration: e.Unix(),
		NotBefore:  t.Unix(),
		Id:         "a",
	}
	fmt.Println(token)

	tr := security.TokenResponse{
		AccessToken:  security.GenerateToken(token),
		TokenType:    configs.BEARER,
		ExpiresIn:    configs.EXPIRY,
		RefreshToken: "N/A",
	}

	b, err := json.Marshal(m)

	if err != nil {
		fmt.Sprintf(err.Error())
	}

	security.VerifyToken(tr.AccessToken)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(b))
}

func MockResponse(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}

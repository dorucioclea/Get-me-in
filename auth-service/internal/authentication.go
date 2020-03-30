package internal

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/auth-service/configs"
	request "github.com/ProjectReferral/Get-me-in/pkg/http_lib"
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

	m := map[string]string{"Authorization": req.Header.Get("Authorization")}
	resp, err := request.Get(configs.LOGIN_ENDPOINT, body, m)

	if err != nil {
		e := err.(*request.ErrorString)
		http.Error(w, e.Reason, e.Code)
		return
	}

	if resp.StatusCode != 200 {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}

	token := IssueToken(req)

	b, err := json.Marshal(token)

	if err != nil {
		fmt.Sprintf(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(b))
}

func IssueToken(req *http.Request) security.TokenResponse{
	t := time.Now()
	e := t.Add(configs.EXPIRY * time.Minute)

	token := &security.TokenClaims{
		Issuer:     configs.SERVICE_ID,
		Subject:    configs.SUBJECT,
		Audience:   req.Header.Get("Origin"),
		IssuedAt:   t.Unix(),
		Expiration: e.Unix(),
		NotBefore:  t.Unix(),
		Id:         req.Header.Get("Id"),
	}

	tr := security.TokenResponse{
		AccessToken:  security.GenerateToken(token),
		TokenType:    configs.BEARER,
		ExpiresIn:    configs.EXPIRY,
		RefreshToken: "N/A",
	}

	return tr
}

func MockResponse(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}

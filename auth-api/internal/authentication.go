package internal

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/auth-api/configs"
	request "github.com/ProjectReferral/Get-me-in/pkg/http_lib"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"io/ioutil"
	"net/http"
	"time"
)

//Validates the request as human/robot with recaptcha
//Validates the credentials via a request to the Account-API
//Token is issued as a JSON with an expiry time of 2.5days
//This token will allow the user to access the [/GET,/PATCH,/DELETE] endpoints for the Account-API
func VerifyCredentials(w http.ResponseWriter, req *http.Request) {

	//TODO: reCaptchacheck

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	resp, err := request.Get(configs.LOGIN_ENDPOINT, body, nil)

	if err != nil {
		e := err.(*request.ErrorString)
		http.Error(w, e.Reason, e.Code)
		return
	}

	if resp.StatusCode != 200 {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}

	token := IssueToken(req, configs.EXPIRY, configs.AUTH_AUTHENTICATED)

	b, err := json.Marshal(token)
	if err != nil {
		fmt.Sprintf(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

//A temporary token can be requested for registration
//This token will only allow the user to access the /PUT endpoint for the Account-API
func IssueRegistrationTempToken(w http.ResponseWriter, req *http.Request){
	token := IssueToken(req, configs.TEMP_EXPIRY, configs.AUTH_REGISTER)

	b, err := json.Marshal(token)

	if err != nil {
		fmt.Sprintf(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func IssueToken(req *http.Request, expiry time.Duration, audience string) security.TokenResponse{
	t := time.Now()
	e := t.Add(expiry * time.Minute)

	token := &security.TokenClaims{
		Issuer:     configs.SERVICE_ID,
		Subject:    configs.SUBJECT,
		Audience:   audience,
		IssuedAt:   t.Unix(),
		Expiration: e.Unix(),
		NotBefore:  t.Unix(),
		Id:         req.Header.Get("id"),
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

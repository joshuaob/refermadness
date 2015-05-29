package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/larryprice/refermadness/utils"
)

type AuthenticationController interface {
	Register(*mux.Router)
}

type AuthenticationControllerImpl struct {
	clientID     string
	clientSecret string
	scheme       string

	session      utils.SessionManager
}

func NewAuthenticationController(clientID, clientSecret string, isDevelopment bool, session utils.SessionManager) *AuthenticationControllerImpl {
	scheme := "http"
	if !isDevelopment {
		scheme += "s"
	}
	return &AuthenticationControllerImpl{
		clientID:     clientID,
		clientSecret: clientSecret,
		scheme:       scheme,
		session:      session,
	}
}

func (ac *AuthenticationControllerImpl) Register(router *mux.Router) {
	router.HandleFunc("/login", ac.login)
	router.HandleFunc("/oauth2callback", ac.oauth2)
}

func (ac *AuthenticationControllerImpl) login(w http.ResponseWriter, r *http.Request) {
	if returnURL := r.FormValue("returnURL"); returnURL != "" {
		ac.session.Set(r, "RedirectAfterLogin", returnURL)
	}

	http.Redirect(w, r, "https://accounts.google.com/o/oauth2/auth?scope=email&redirect_uri="+
		ac.scheme+"%3A%2F%2F"+r.Host+"%2foauth2callback"+"&response_type=code&client_id="+ac.clientID,
		http.StatusTemporaryRedirect)
}

func (ac *AuthenticationControllerImpl) oauth2(w http.ResponseWriter, r *http.Request) {
	redirectTo := "/"
	if returnURL := ac.session.Get(r, "RedirectAfterLogin"); returnURL != "" {
		ac.session.Delete(r, "RedirectAfterLogin")
		redirectTo = returnURL
	}

	if r.FormValue("error") != "" {
		fmt.Println("Error in OAuth", r.FormValue("error"))
		http.Redirect(w, r, redirectTo, http.StatusFound)
		return
	}
	// send token request
	resp, err := http.PostForm("https://www.googleapis.com/oauth2/v3/token",
		url.Values{"code": {r.FormValue("code")}, "grant_type": {"authorization_code"}, "redirect_uri": {ac.scheme + "://" + r.Host + "/oauth2callback"},
			"client_id": {ac.clientID}, "client_secret": {ac.clientSecret}})

	var result map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	token, _ := jwt.Parse(result["id_token"].(string), func(token *jwt.Token) (interface{}, error) {
		return result["access_token"], nil
	})
	fmt.Println(token.Claims["email"], err, resp.StatusCode)

	http.Redirect(w, r, redirectTo, http.StatusFound)
}

package routes

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
)

func (c *Routes) register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if r.Header.Get("Content-Type") != "application/json" {
		jsonError(w, http.StatusBadRequest, "unknown payload")
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	_ = r.Body.Close()

	fd := &struct {
		Login    string `json:"username"`
		Password string `json:"password"`
	}{}
	err := json.Unmarshal(body, fd)
	if err != nil {
		jsonError(w, http.StatusBadRequest, "cant unpack payload")
		return
	}

	user, err := c.users.Add(r.Context(), fd.Login, fd.Password)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "bad login or password")
		return
	}

	sessionID, err := c.session.Write(r.Context(), user.ID)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "unable to save session data")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": jwt.MapClaims{
			"username": user.Login,
			"id":       user.ID,
		},
		"sessionID": sessionID,
	})
	tokenString, err := token.SignedString(c.authSecret)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, _ := json.Marshal(map[string]interface{}{
		"token": tokenString,
	})
	_, _ = w.Write(resp)
	_, _ = w.Write([]byte("\n\n"))
}

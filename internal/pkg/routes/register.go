package routes

import (
	"encoding/json"
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

	tokenString, err := c.auth.BeginSession(r.Context(), user.ID, user.Login)
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

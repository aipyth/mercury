package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mercuryapi/db"
	"net/http"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenOnly struct {
	Token string `json:"token"`
}

var defaultContext = context.Background()

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var qs = r.Context().Value("DB").(*db.Queries)
	lg := r.Context().Value("logging").(Logging)

	// read request body
	userInfo := UserInfo{}
	err := UnpackBody(r.Body, &userInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"errors":["Error during decoding request body."]}`))
		return
	}

	// send BadRequest if email or password are not specified
	if userInfo.Email == "" || userInfo.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"errors":["Empty email or password."]}`))
		return
	}

	// save user
	user := db.CreateUserParams{
		Email:    userInfo.Email,
		Password: HashPassword(userInfo.Password),
		Token:    uuid.New().String(),
	}

	err = qs.CreateUser(defaultContext, user)
	if e, ok := err.(*pq.Error); ok && e.Code.Name() == "unique_violation" {
		// have psql error, so this email address already exists
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"errors":["This user already exists."]}`))
		return
	} else if err != nil {
		// some other error
		lg.Error(err)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"errors":["User creation failed."]}`))
		return
	}

	// process response
	userInfo.Password = ""
	userBytes, err := json.Marshal(userInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"errors\":[\"Error during encoding response json." +
			"\"]}"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(userBytes)
}

func ObtainUserTokenHandler(w http.ResponseWriter, r *http.Request) {
	var qs = r.Context().Value("DB").(*db.Queries)
	//lg := r.Context().Value("logging").(Logging)

	// read request body
	userInfo := UserInfo{}
	err := UnpackBody(r.Body, &userInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"errors":["Error during decoding request body."]}`))
		return
	}

	user, err := qs.GetUser(defaultContext, userInfo.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"errors":["Error during getting user object. Check user credentials."]}`))
		return
	}

	if HashPassword(userInfo.Password) == user.Password {
		token := TokenOnly{Token: user.Token}
		tokenBytes, _ := json.Marshal(token)
		w.WriteHeader(http.StatusAccepted)
		w.Write(tokenBytes)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var qs = r.Context().Value("DB").(*db.Queries)
		//lg := r.Context().Value("logging").(Logging)

		data, _ := ioutil.ReadAll(r.Body)

		dataJson := TokenOnly{}
		json.Unmarshal(data, &dataJson)

		var user db.User
		if dataJson.Token != "" {
			user, _ = qs.GetUserByToken(defaultContext, dataJson.Token)
			//ctx := context.WithValue(r.Context(), "user", user)
		} else {
			//ctx := context.WithValue(r.Context(), "user", user)
		}
		fmt.Println("in handler", user)
		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
		r.Body = ioutil.NopCloser(bytes.NewReader(data))
		h.ServeHTTP(w, r)
	})
}

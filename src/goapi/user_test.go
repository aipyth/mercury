package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mercuryapi/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func TestCreateUserHandler(t *testing.T) {
	tt := []TestObj{
		{
			name:   "right request",
			method: http.MethodPost,
			input: &UserInfo{Email: "example@mail.com",
				Password: "12345678"},
			want:       `{"email":"example@mail.com","password\":""}`,
			statusCode: http.StatusCreated,
		},
		{
			name:   "unique violation test",
			method: http.MethodPost,
			input: &UserInfo{Email: "example@mail.com",
				Password: "12345678"},
			want:       `{"errors":\["This user already exists\."\]}`,
			statusCode: http.StatusConflict,
		},
		{
			name:       "empty email",
			method:     http.MethodPost,
			input:      &UserInfo{Password: "12345678"},
			want:       `{"errors":\["Empty email or password\."\]}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "empty password",
			method:     http.MethodPost,
			input:      &UserInfo{Email: "example@mail.com"},
			want:       `{"errors":\["Empty email or password\."\]}`,
			statusCode: http.StatusBadRequest,
		},
	}

	d := GetDBConnection()
	StandardTest(t, tt, CreateUserHandler, d)
}

func TestObtainUserTokenHandler(t *testing.T) {
	userPssw := "12345678"
	token := uuid.New().String()
	user := db.CreateUserParams{
		Email:    "example@mail.com",
		Password: HashPassword(userPssw),
		Token:    token,
	}
	tt := []TestObj{
		{
			name:   "right request",
			method: http.MethodPost,
			input: &UserInfo{Email: user.Email,
				Password: userPssw},
			want:       "^{\"token\":\"(.+)\"}$",
			statusCode: http.StatusAccepted,
		},
		{
			name:   "bad credentials",
			method: http.MethodPost,
			input: &UserInfo{Email: user.Email + "o",
				Password: userPssw},
			want: `{"errors":\["Error during getting user object\. ` +
				`Check user credentials\."\]}`,
			statusCode: http.StatusBadRequest,
		},
	}

	d := GetDBConnection()
	err := d.CreateUser(context.Background(), user)
	if e, ok := err.(*pq.Error); ok && e.Code.Name() == "unique_violation" {
		log.Println(err)
	} else if err != nil {
		log.Println(err)
	}
	StandardTest(t, tt, ObtainUserTokenHandler, d)
}

func TestAuthMiddleware(t *testing.T) {
	d := GetDBConnection()

	userPssw := "12345678"
	token := uuid.New().String()
	user := db.CreateUserParams{
		Email:    "example@mail.com",
		Password: HashPassword(userPssw),
		Token:    token,
	}
	err := d.CreateUser(context.Background(), user)
	if e, ok := err.(*pq.Error); ok && e.Code.Name() == "unique_violation" {
		log.Println(err)
	} else if err != nil {
		log.Println(err)
	}
	u, _ := d.GetUser(defaultContext, user.Email)
	token = u.Token

	tt := []struct {
		name       string
		method     string
		input      string
		userExists bool
	}{
		{
			name:       "normal request",
			method:     http.MethodPost,
			input:      "{\"token\":\"" + token + "\"}",
			userExists: true,
		},
		{
			name:       "no user request",
			input:      "{}",
			userExists: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			lg := Logging{Level: 1}

			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				u, ok := r.Context().Value("user").(db.User)
				if !ok {
					t.Error("user is not of type db.User")
				} else if tc.userExists != (u.Email != "") {
					fmt.Println(u)
					t.Error("user instance is (not) present in next request")
				}
			})

			request := httptest.NewRequest(tc.method, "/api/user/token",
				bytes.NewReader([]byte(tc.input)))
			responseRecorder := httptest.NewRecorder()

			ctx := request.Context()
			ctx = context.WithValue(ctx, "DB", d)
			ctx = context.WithValue(ctx, "logging", lg)

			request = request.WithContext(ctx)

			handler := AuthMiddleware(nextHandler)
			handler.ServeHTTP(responseRecorder, request)

		})
	}
}

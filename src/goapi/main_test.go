package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"mercuryapi/db"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
)

type TestObj struct {
	name       string
	method     string
	input      interface{}
	want       string
	statusCode int
}

func GetDBConnection() *db.Queries {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s" +
		" port=%s sslmode=disable" +
		" TimeZone=Europe/Kiev", os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	q, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	d := db.New(q)
	return d
}

func StandardTest(t *testing.T, tt []TestObj, f func(w http.ResponseWriter,
	r *http.Request), d *db.Queries) {
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			lg := Logging{Level: 1}

			body, _ := json.Marshal(tc.input)

			request := httptest.NewRequest(tc.method, "/api/user",
				bytes.NewReader(body))
			responseRecorder := httptest.NewRecorder()

			ctx := request.Context()
			ctx = context.WithValue(ctx, "DB", d)
			ctx = context.WithValue(ctx, "logging", lg)

			request = request.WithContext(ctx)

			handler := http.HandlerFunc(f)
			handler.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode,
					responseRecorder.Code)
			}

			re := regexp.MustCompile(tc.want)
			//fmt.Println(re.Match([]byte(`seafood fool`)))

			//if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
			if !re.Match([]byte(strings.TrimSpace(responseRecorder.Body.
				String()))) {
				t.Errorf("Want: \n'%s', got: \n'%s'", tc.want,
					responseRecorder.Body)
			}
		})
	}
}

package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"log"
	"mercuryapi/db"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	lg := Logging{Level: 1}

	// run db start file
	dbStartFilePath := os.Getenv("DB_START_FILE_PATH")
	cmd := exec.Command("psql", "-U", os.Getenv("DB_USER"),
		"-h", os.Getenv("DB_HOST"), "-d", os.Getenv("DB_NAME"), "-a",
		"-f", dbStartFilePath)

	var out, stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	}


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s" +
		" port=%s sslmode=disable" +
		" TimeZone=Europe/Kiev", os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	q, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer q.Close()
	qs := db.New(q)



	attachToContext := func(name string, value interface{}) func (
			http.Handler) http.Handler{
		return func (handler http.Handler) http.Handler {
			return http.HandlerFunc(func(
					w http.ResponseWriter,
					r *http.Request) {
				ctx := context.WithValue(r.Context(), name, value)
				handler.ServeHTTP(w, r.WithContext(ctx))
			})
		}
	}


	r := chi.NewRouter()

	r.Use(middleware.Logger)
	// give db, logger to other handlers through context
	r.Use(attachToContext("DB", qs))
	r.Use(attachToContext("logging", lg))
	r.Use(AuthMiddleware)



	r.Route("/api", func (r chi.Router){
		r.Route("/timeschemas", func (r chi.Router){
			r.Post("/", CreateTimeSchemaHandler)
			r.Get("/", RetrieveTimeSchemaHandler)
			r.Put("/", UpdateTimeSchemaHandler)
		})
		//r.Route("/rooms", func (r chi.Router){
		//	r.Post("/", CreateRoomHandler)
		//})
		//r.Route("/subjects", func (r chi.Router){
		//	//r.Post("/", CreateSubjectHandler)
		//})
		r.Route("/user", func (r chi.Router){
			r.Post("/", CreateUserHandler)
			r.Post("/token", ObtainUserTokenHandler)
		})
	})
	http.ListenAndServe(":8080", r)
}

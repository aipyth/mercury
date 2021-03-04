package apiutils

import (
	"net/http"
	"log"
)

func Fob() {
	log.Println("FOOOB")
}

func GetPublicRooms() {
	response, err := http.Get("http://web:8000/api/timeshemes/")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
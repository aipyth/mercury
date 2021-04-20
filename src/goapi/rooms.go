package main

import (
	"fmt"
	"net/http"
	"time"
)

type CreateRoomInfo struct {
	Name       string      `json:"name"`
	Slug       string      `json:"slug"`
	Period     int64       `json:"period"`
	StartDate  time.Time   `json:"start-date"`
	EndDate    time.Time   `json:"end-date"`
	Public     bool        `json:"public"`
	TimeSchema interface{} `json:"time-schema"`
}

type ResponseRoomInfo struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Slug      string
	Period    int64
	StartDate time.Time
	EndDate   time.Time
	Public    bool
}

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	//db := r.Context().Value("DB").(*gorm.DB)
	//lg := r.Context().Value("logging").(	Logging)

	roomInfo := CreateRoomInfo{}
	err := UnpackBody(r.Body, &roomInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"errors\":[\"Error decoding json.\"]}"))
	}

	// check if TimeSchema object is valid in request
	timeSchemaInfo := UpdateTimeSchemaInfo{}
	// is TimeSchema integer?
	v, ok := roomInfo.TimeSchema.(uint)
	if ok {
		// ok! get the object from db
		timeSchemaInfo.Id = v
		fmt.Println(timeSchemaInfo)
	} else {
		// or its an object
		timeSchemaInfo, ok := roomInfo.TimeSchema.(UpdateTimeSchemaInfo)
		fmt.Println(ok, timeSchemaInfo)
	}
	fmt.Println(roomInfo)

	// TODO check if all required fields are present

	//room := &orm.Room{
	//	Name: roomInfo.Name,
	//	Slug: roomInfo.Slug,
	//	Period: roomInfo.Period,
	//	StartDate: roomInfo.StartDate,
	//	EndDate: roomInfo.EndDate,
	//	Public: roomInfo.Public,
	//}
	//
	//// check if timeschema in request exists
	//var tsCount int64
	//db.Model(&orm.TimeSchema{}).Where(&orm.
	//TimeSchema{ID: }).Count(&tsCount)
	//if tsCount == 0 {
	//	w.WriteHeader(http.StatusBadRequest)
	//	w.Write([]byte("{\"errors\":[\"No such object.\"]}"))
	//	return
	//}

}

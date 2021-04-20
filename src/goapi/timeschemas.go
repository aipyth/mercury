package main

import (
	"encoding/json"
	"mercuryapi/db"
	"net/http"
	"time"

	"github.com/lib/pq"
)

type CreateTimeSchemaInfo struct {
	Name  string
	Items interface{}
}

type UpdateTimeSchemaInfo struct {
	Id uint
	CreateTimeSchemaInfo
}

type RetrieveTimeSchemaInfo struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     interface{}
}

func CreateTimeSchemaHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.Context().Value("DB").(*db.Queries)
	lg := r.Context().Value("logging").(Logging)

	timeSchemaData := CreateTimeSchemaInfo{}
	UnpackBody(r.Body, &timeSchemaData)
	// encode timeschema items back to json
	items, _ := json.Marshal(timeSchemaData.Items)

	timeSchema := db.CreateTimeSchemaParams{
		Name:  timeSchemaData.Name,
		Items: items,
	}

	ts, err := qs.CreateTimeSchema(defaultContext, timeSchema)
	if e, ok := err.(*pq.Error); ok && e.Code.Name() == "unique_violation" {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"errors":["This timeschema already exists."]}`))
		return
	} else if err != nil {
		lg.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"errors":["TimeSchema creation failed."]}`))
		return
	}
	tsb, _ := json.Marshal(ts)
	w.WriteHeader(http.StatusCreated)
	w.Write(tsb)
}

func RetrieveTimeSchemaHandler(w http.ResponseWriter, r *http.Request) {
	//db := r.Context().Value("DB").(*db.Queries)
	////lg := r.Context().Value("logging").(Logging)
	//
	//requestData := struct {
	//	Id 	 uint
	//	Name string
	//}{}
	//DecodeBody(r.Body, &requestData)
	//
	//var timeSchemasObjs []orm.TimeSchema
	//
	//if requestData.Id != 0 {
	//	// request by id
	//	db.Model(&orm.TimeSchema{}).Where("id = ? ",
	//		requestData.Id).Find(&timeSchemasObjs)
	//} else if requestData.Name != "" {
	//	// request by name
	//	db.Model(&orm.TimeSchema{}).Where("name LIKE ?",
	//		requestData.Name+"%").Find(&timeSchemasObjs)
	//} else {
	//	// nothing is specified so return bad request
	//	w.WriteHeader(http.StatusBadRequest)
	//	w.Write([]byte("{\"errors\":[\"Neither id nor name specified.\"]}"))
	//	return
	//}
	//
	//timeSchemas := make([]RetrieveTimeSchemaInfo, len(timeSchemasObjs))
	//
	//for i, v := range timeSchemasObjs {
	//	timeSchemas[i].Id = v.ID
	//	timeSchemas[i].Name = v.Name
	//	timeSchemas[i].CreatedAt = v.CreatedAt
	//	timeSchemas[i].UpdatedAt = v.UpdatedAt
	//	json.Unmarshal(v.Items, &timeSchemas[i].Items)
	//}
	//
	//tsb, err := json.Marshal(timeSchemas)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	w.Write([]byte("{\"errors\":[\"Error encoding response\"]}"))
	//	return
	//}
	//w.WriteHeader(http.StatusOK)
	//w.Write(tsb)
}

func UpdateTimeSchemaHandler(w http.ResponseWriter, r *http.Request) {
	//db := r.Context().Value("DB").(*gorm.DB)
	//lg := r.Context().Value("logging").(Logging)

}

func SearchTimeSchemaHandler(w http.ResponseWriter, r *http.Request) {

}

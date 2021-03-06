package dbutils

import (
	"log"
	"database/sql"
	"encoding/json"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func rowsToRooms(rows *sql.Rows) []Room {
	var rooms []Room = make([]Room, 0)
	var i = 0
	for rows.Next() {
		rooms = append(rooms, Room{})
		if err := rows.Scan(
				&rooms[i].Id, &rooms[i].Name, &rooms[i].Slug,
				&rooms[i].Period, &rooms[i].StartDate,
				&rooms[i].EndDate, &rooms[i].Public,
				&rooms[i].Created, &rooms[i].OwnerId,
				&rooms[i].TimeSchemaId, &rooms[i].ScheduleImage,
				&rooms[i].ScheduleImageThumb);
			err != nil {
			log.Fatal(err)
		}
		rooms[i].TimeSchema = rooms[i].GetTimeSchema()
		i++
	}
	return rooms
}

func GetRoom(slug string) *Room {
	rows, err := DB.Query(`SELECT id, name, slug, period, start_date,
								  end_date, public, created,
								  owner_id, time_schema_id,
								  schedule_image, schedule_image_thumb
						   FROM schedule_room
						   WHERE slug=$1
						   ;`, slug)
	if err != nil {
		log.Fatal(err)
	}
	rooms := rowsToRooms(rows)
	if len(rooms) == 0 {
		return nil
	}
	return &rooms[0]
}

func GetRoomsBySlug(slug string, offset uint64) []Room {
	pattern := slug + "%"
	rows, err := DB.Query(`SELECT id, name, slug, period, start_date,
								  end_date, public, created,
								  owner_id, time_schema_id,
								  schedule_image, schedule_image_thumb
						   FROM schedule_room
						   WHERE slug LIKE $1
						   ORDER BY substr(slug, 0)
						   LIMIT $2
						   OFFSET $3
						   ;`, pattern, INLINE_PAGINATION_LIMIT, offset)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return rowsToRooms(rows)	
}

func rowsToSubjects(rows *sql.Rows) []Subject {
	var subjects []Subject = make([]Subject, 0)
	var i = 0
	for rows.Next() {
		subjects = append(subjects, Subject{})
		if err := rows.Scan(
				&subjects[i].Id, &subjects[i].Name,
				&subjects[i].daysAndOrders, &subjects[i].Lector,
				&subjects[i].Extra, &subjects[i].RoomId);
			err != nil {
			log.Fatal(err)
		}
		json.Unmarshal([]byte(subjects[i].daysAndOrders),
			&(subjects[i].DaysAndOrders))
		i++
	}
	return subjects
}


func (r *Room) GetSubjects() []Subject {
	rows, err := DB.Query(`SELECT id, name, days_and_orders, lector,
								  extra, room_id
						   FROM schedule_subject
						   WHERE room_id=$1`, r.Id)
	if err != nil {
		log.Print("Error getting subjects. ", err)
	}
	defer rows.Close()
	return rowsToSubjects(rows)
}

func rowsToTimeSchemas(rows *sql.Rows) []TimeSchema {
	var ts []TimeSchema = make([]TimeSchema, 0)
	var i = 0
	for rows.Next() {
		ts = append(ts, TimeSchema{})
		if err := rows.Scan(
				&ts[i].Id, &ts[i].Name, &ts[i].items,
				&ts[i].Public, &ts[i].Created);
			err != nil {
			log.Fatal(err)
		}
		json.Unmarshal([]byte(ts[i].items),
			&(ts[i].Items))
		i++
	}
	return ts
}

func (r *Room) GetTimeSchema() TimeSchema {
	// TODO: optimize for one record in rows
	rows, err := DB.Query(`SELECT id, name, items, public, created
						   FROM schedule_timeschema
						   WHERE id=$1`, r.TimeSchemaId)
	if err != nil {
		log.Print("Error getting TimeSchema. ", err)
	}
	defer rows.Close()
	return rowsToTimeSchemas(rows)[0]
}
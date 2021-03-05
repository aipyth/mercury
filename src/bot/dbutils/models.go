package dbutils

import (
	"time"
	"fmt"
	"math"
)

type TimeSchema struct {
	Id int64
	Name string
	items string
	Public bool
	Created time.Time

	Items []map[string]string
}

type Room struct {
	Id int64
	Name string
	Slug string
	Period int64
	StartDate time.Time
	EndDate time.Time
	Public bool
	Created time.Time
	OwnerId int64
	TimeSchemaId int64
	ScheduleImage string
	ScheduleImageThumb string

	TimeSchema TimeSchema
}

type Subject struct {
	Id int64
	Name string
	daysAndOrders string
	Lector string
	Extra string
	RoomId int64

	DaysAndOrders map[int64][]int64
}

type DaySchedule struct {
	RoomInstance *Room

	Day int64
	Weekday time.Weekday
	Subjects map[int64]Subject
}

func (t* TimeSchema) Entries() int {
	return len(t.Items)
}

func (t* TimeSchema) TimeToRepresentation() []string {
	r := make([]string, 0)
	const timeOnly = "15:04:05"
	for _, m := range t.Items {
		start, _ := time.Parse(timeOnly, m["start"])
		stop, _ := time.Parse(timeOnly, m["stop"])
		r = append(r, fmt.Sprintf("<b>%d:%02d - %d:%02d</b>", start.Hour(), start.Minute(),
			stop.Hour(), stop.Minute()))
	}
	return r
}

func (r *Room) DayToday() (int64, time.Weekday) {
	offset := time.Now().Sub(r.StartDate)
	day := int64(math.Floor(offset.Hours() / 24)) % r.Period
	return day, time.Now().Weekday()
}

func (r *Room) DayTomorrow() (int64, time.Weekday) {
	offset := time.Now().Sub(r.StartDate)
	day := int64(math.Floor(offset.Hours() / 24) + 1) % r.Period
	return day, time.Now().Weekday() + 1
}

func (r *Room) ScheduleToday() DaySchedule {
	subjects := r.GetSubjects()
	day, weekday := r.DayToday()
	schedule := DaySchedule{
		RoomInstance: r,
		Day: day,
		Weekday: weekday,
		Subjects: make(map[int64]Subject),
	}

	for _, subject := range subjects {
		if len(subject.DaysAndOrders[day]) != 0 {
			for _, order := range subject.DaysAndOrders[day] {
				schedule.Subjects[order] = subject
			}
		}
	}
	return schedule
}

func (r *Room) ScheduleTomorrow() DaySchedule {
	subjects := r.GetSubjects()
	day, weekday := r.DayTomorrow()
	schedule := DaySchedule{
		RoomInstance: r,
		Day: day,
		Weekday: weekday,
		Subjects: make(map[int64]Subject),
	}

	for _, subject := range subjects {
		if len(subject.DaysAndOrders[day]) != 0 {
			for _, order := range subject.DaysAndOrders[day] {
				schedule.Subjects[order] = subject
			}
		}
	}
	return schedule
}

func (s Subject) ToRepresentation() string {
	if s.Extra == "" {
		return fmt.Sprintf("<i>%s</i>\n    %s", s.Name, s.Lector)	
	}
	return fmt.Sprintf("<i>%s</i>\n    %s\n%s", s.Name, s.Lector, s.Extra)
}

func (d DaySchedule) ToRepresentation() string {
	output := fmt.Sprintf("\t<u><b>%s</b></u> %s\n",
		d.RoomInstance.Name, d.Weekday)

	add := func (s ...string) {
		for _, ss := range s {
			output = output + ss
		}
	}
	timeSchemaReprs := d.RoomInstance.TimeSchema.TimeToRepresentation()
	for i := 0; i < d.RoomInstance.TimeSchema.Entries(); i++ {
		add(fmt.Sprintf("\n<b>%d.    </b> ", i+1))
		if d.Subjects[int64(i)].DaysAndOrders != nil {
			add(timeSchemaReprs[i])
			add("\n   ")
			add(d.Subjects[int64(i)].ToRepresentation())
			add("\n")
		}
	}

	return output
}
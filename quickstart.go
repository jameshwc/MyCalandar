package mycalendar

import (
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
)

func main() {
	srv := initCalendar()
	event := &calendar.Event{
		Summary:     "Test",
		Location:    "NTU",
		Description: "A chance to hear more about Google's developer products.",
		Start: &calendar.EventDateTime{
			DateTime: "2020-02-08T09:00:00-07:00",
			TimeZone: "Asia/Taipei",
		},
		End: &calendar.EventDateTime{
			DateTime: "2020-02-08T17:00:00-07:00",
			TimeZone: "Asia/Taipei",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		// Attendees: []*calendar.EventAttendee{
		// 	&calendar.EventAttendee{Email: "lpage@example.com"},
		// 	&calendar.EventAttendee{Email: "sbrin@example.com"},
		// },
	}

	calendarID := "primary"
	event, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
	fmt.Printf("%T %T %T", event, calendarID, srv)

}

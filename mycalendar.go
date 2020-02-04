package mycalendar

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("There should be at least one argument!")
	}
	srv := initCalendar()
	switch os.Args[1] {
	case "print":
		listUpcomingEvents(srv)
	case "add":

	}
}

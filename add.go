package mycalendar

import (
	"bufio"
	"fmt"
	"os"

	"google.golang.org/api/calendar/v3"
)

func input() calendar.Event {
	var title, location, description string
	var startDate, endDate, startTime, endTime string
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Event Title:")
	fmt.Scanln(&title)
	fmt.Println("Location:")
	fmt.Scanln(&location)
	fmt.Println("Description:")
	fmt.Scanln(&description)
	fmt.Println("Is this event ")
	fmt.Println("Start date (format:MM/DD):")
	fmt.


}
func add(srv *calendar.Service) {
	event := input()
}

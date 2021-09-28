package main

import (
	"log"
	"time"

	"github.com/caius/event-sandbox/internal/events"
)

type userCreatedNotifier struct{}

func (u userCreatedNotifier) notifyAdmin(email string, time time.Time) {
	log.Printf("ADMIN ALERT: %s created at %v\n", email, time)
}

func (u userCreatedNotifier) Handle(payload events.UserCreatedPayload) {
	u.notifyAdmin(payload.Email, payload.Time)
}

type userDeletedNotifier struct{}

func (u userDeletedNotifier) notifyAdmin(email string, time time.Time) {
	log.Printf("ADMIN ALERT: %s deleted at %v\n", email, time)
}

func (u userDeletedNotifier) Handle(payload events.UserDeletedPayload) {
	u.notifyAdmin(payload.Email, payload.Time)
}

func main() {
	createNotifier := userCreatedNotifier{}
	events.UserCreated.Register(createNotifier)

	deleteNotifier := userDeletedNotifier{}
	events.UserDeleted.Register(deleteNotifier)

	log.Printf("Start")

	events.UserCreated.Trigger(events.UserCreatedPayload{
		Email: "dickhead@example.com",
		Time:  time.Now(),
	})

	events.UserDeleted.Trigger(events.UserDeletedPayload{
		Email: "dickhead@example.com",
		Time:  time.Now(),
	})

	time.Sleep(1 * time.Second)
	log.Printf("And done")
}

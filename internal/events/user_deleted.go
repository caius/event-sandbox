package events

import (
	"time"
)

var UserDeleted userDeleted

type UserDeletedPayload struct {
	Email string
	Time  time.Time
}

type userDeleted struct {
	handlers []interface{ Handle(UserDeletedPayload) }
}

// Register adds an event handler for this event
func (u *userDeleted) Register(handler interface{ Handle(UserDeletedPayload) }) {
	u.handlers = append(u.handlers, handler)
}

// Trigger sends out an event with the payload
func (u userDeleted) Trigger(payload UserDeletedPayload) {
	for _, handler := range u.handlers {
		go handler.Handle(payload)
	}
}

package types

import "time"

type EventPayload struct {
  Title string `json:"title"`
  Description string `json:"description"`
  Public bool `json "public"`
  Date time.Time `json:"date"`
}


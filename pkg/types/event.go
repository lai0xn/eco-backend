package types

import "time"

type EventPayload struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Public      bool      `json:"public"`
	Date        time.Time `json:"date"`
	OrgID       string    `json:"orgId"`
	Location    string    `json:"location"`
}

type AcheivmentPayload struct {
	Title   string `json:"string"`
	Details string `json:"details"`
	EventID string `json:"eventId"`
	OrgID   string `json:"orgId"`
}

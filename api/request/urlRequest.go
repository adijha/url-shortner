package request

import "time"

type Url struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
	User_id     string        `json:"user_id"`
}
